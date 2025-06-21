use crate::genurl::common::Url;
use crate::{genurl::common::Assets, types::FulfillmentDataContent};
use anyhow::{Result, anyhow};
use futures::future;
use reqwest::Client;
use roxmltree::Document;
use std::collections::HashMap;
use std::sync::Arc;
use std::time::Duration;
use tokio::sync::Semaphore;
use tokio::time::timeout;

/// Generates UWP download URLs using the Microsoft Store delivery API.
///
/// # Arguments
/// * `client` - A reqwest client instance to make HTTP requests
/// * `fulfillment_data` - JSON string containing product information including WuCategoryId
///
/// # Returns
/// * `Ok(HashMap<String, String>)` - Map of filenames to download URLs
/// * `Err(anyhow::Error)` - If any step in the process fails
///
pub async fn gen_uwp_url(client: &Client, fulfillment_data: &str) -> Result<Url> {
    // 1. Parse input data
    let list: FulfillmentDataContent = serde_json::from_str(fulfillment_data)
        .map_err(|e| anyhow!("Failed to parse fulfillment data: {}", e))?;

    let category_id = list.wu_category_id;
    const RELEASE_TYPE: &str = "retail";

    // 2. Get encrypted cookie
    let cookie_template_asset = Assets::get("GetCookie.xml").ok_or_else(|| {
        anyhow!(
            "Failed to get cookie template 'GetCookie.xml'. Make sure it's in the 'assets/' folder."
        )
    })?;

    let resp = client
        .post("https://fe3cr.delivery.mp.microsoft.com/ClientWebService/client.asmx")
        .header("Content-Type", "application/soap+xml; charset=utf-8")
        .body(cookie_template_asset.data.into_owned())
        .send()
        .await
        .map_err(|e| anyhow!("Failed to send cookie request: {}", e))?;

    let resp_text = resp
        .text()
        .await
        .map_err(|e| anyhow!("Failed to get cookie response text: {}", e))?;

    let cookie_doc = Document::parse(&resp_text)
        .map_err(|e| anyhow!("Failed to parse cookie response XML: {}", e))?;

    let encrypted_data = cookie_doc
        .root_element()
        .descendants()
        .find(|n| n.tag_name().name() == "EncryptedData")
        .and_then(|n| n.text())
        .ok_or_else(|| anyhow!("Failed to find 'EncryptedData' in cookie response."))?;

    // 3. Request ID and filenames
    let wuid_template_asset = Assets::get("WUIDRequest.xml").ok_or_else(|| {
        anyhow!(
            "Failed to get WUID template 'WUIDRequest.xml'. Make sure it's in the 'assets/' folder."
        )
    })?;

    let wuid_template = String::from_utf8_lossy(&wuid_template_asset.data)
        .replace("{0}", encrypted_data)
        .replace("{1}", &category_id)
        .replace("{2}", RELEASE_TYPE);

    let resp = client
        .post("https://fe3cr.delivery.mp.microsoft.com/ClientWebService/client.asmx")
        .header("Content-Type", "application/soap+xml; charset=utf-8")
        .body(wuid_template)
        .send()
        .await
        .map_err(|e| anyhow!("Failed to send WUID request: {}", e))?;

    let resp_text = resp
        .text()
        .await
        .map_err(|e| anyhow!("Failed to get WUID response text: {}", e))?;
    let decoded_resp_text = html_escape::decode_html_entities(&resp_text);
    let xml_doc = Document::parse(&decoded_resp_text)
        .map_err(|e| anyhow!("Failed to parse WUID response XML: {}", e))?;

    let mut filenames_map: HashMap<String, (String, String)> = HashMap::new();
    for files_node in xml_doc
        .descendants()
        .filter(|n| n.tag_name().name() == "Files")
    {
        if let Some(package_container_node) = files_node.parent().and_then(|p| p.parent()) {
            if let Some(package_id_node) = package_container_node
                .descendants()
                .find(|n| n.tag_name().name() == "ID")
            {
                if let Some(node_id) = package_id_node.text() {
                    if let Some(file_node_in_files) = files_node
                        .descendants()
                        .find(|n| n.tag_name().name() == "File")
                    {
                        let installer_prefix = file_node_in_files
                            .attribute("InstallerSpecificIdentifier")
                            .unwrap_or_default()
                            .to_string();
                        let fname = file_node_in_files
                            .attribute("FileName")
                            .unwrap_or_default()
                            .to_string();
                        let modified = file_node_in_files
                            .attribute("Modified")
                            .unwrap_or_default()
                            .to_string();
                        filenames_map.insert(
                            node_id.to_string(),
                            (format!("{}_{}", installer_prefix, fname), modified),
                        );
                    }
                }
            }
        }
    }

    if filenames_map.is_empty() {
        return Err(anyhow!(
            "Server returned an empty list of files from WUID request."
        ));
    }

    // 4. Parse update ID from SecuredFragment
    let mut identities: HashMap<String, (String, String)> = HashMap::new();

    for fragment_node in xml_doc
        .descendants()
        .filter(|n| n.tag_name().name() == "SecuredFragment")
    {
        if let Some(great_grandparent) = fragment_node.ancestors().nth(3) {
            if let Some(package_id_node) = great_grandparent
                .descendants()
                .find(|n| n.tag_name().name() == "ID")
            {
                if let Some(fn_id_text) = package_id_node.text() {
                    let fn_id = fn_id_text.to_string();

                    if let Some((prefixed_filename, _modified_date)) = filenames_map.get(&fn_id) {
                        if let Some(grandparent) = fragment_node.ancestors().nth(2) {
                            if let Some(top_node) = grandparent.first_child() {
                                let update_id = top_node
                                    .attribute("UpdateID")
                                    .unwrap_or_default()
                                    .to_string();
                                let rev_num = top_node
                                    .attribute("RevisionNumber")
                                    .unwrap_or_default()
                                    .to_string();
                                identities.insert(prefixed_filename.clone(), (update_id, rev_num));
                            }
                        }
                    }
                }
            }
        }
    }

    // 5. Get download URLs for all content
    let semaphore = Arc::new(Semaphore::new(10));
    let mut tasks = Vec::new();

    let file_url_template_asset = Assets::get("FE3FileUrl.xml")
        .ok_or_else(|| anyhow!("Failed to get file URL template 'FE3FileUrl.xml'. Make sure it's in the 'assets/' folder."))?;
    let file_url_template = String::from_utf8_lossy(&file_url_template_asset.data).to_string();

    let base_client = client.clone();

    for (file_name, (update_id, revision_num)) in identities.clone().into_iter() {
        let client_clone = base_client.clone();
        let file_url_template_clone = file_url_template.clone();
        let release_type_clone = RELEASE_TYPE.to_string();
        let semaphore_clone = semaphore.clone();

        tasks.push(tokio::spawn(async move {
            let permit = semaphore_clone.acquire_owned().await.expect("Failed to acquire semaphore permit");

            let request_body = file_url_template_clone
                .replace("{0}", &update_id)
                .replace("{1}", &revision_num)
                .replace("{2}", &release_type_clone);

            let resp_result = timeout(Duration::from_secs(15), client_clone
                .post("https://fe3cr.delivery.mp.microsoft.com/ClientWebService/client.asmx/secured")
                .header("Content-Type", "application/soap+xml; charset=utf-8")
                .body(request_body)
                .send()
            ).await;

            drop(permit);

            let resp = match resp_result {
                Ok(Ok(r)) => r,
                Ok(Err(e)) => {
                    eprintln!("Error sending URL request for '{}': {}", file_name, e);
                    return None;
                },
                Err(_) => {
                    eprintln!("URL request timed out for '{}'", file_name);
                    return None;
                }
            };

            let resp_text = match resp.text().await {
                Ok(t) => t,
                Err(e) => {
                    eprintln!("Error getting response text for '{}': {}", file_name, e);
                    return None;
                }
            };

            let doc = match Document::parse(&resp_text) {
                Ok(d) => d,
                Err(e) => {
                    eprintln!("Error parsing URL response XML for '{}': {}", file_name, e);
                    return None;
                }
            };

            if let Some(file_location_node) = doc.descendants().find(|n| n.tag_name().name() == "FileLocation") {
                if let Some(url_node) = file_location_node.descendants().find(|n| n.tag_name().name() == "Url") {
                    if let Some(url) = url_node.text() {
                        return Some((file_name, url.to_string()));
                    }
                }
            }
            None
        }));
    }

    let results = future::join_all(tasks).await;

    let mut file_dict: HashMap<String, String> = HashMap::new();
    for (file_name, url) in results.into_iter().flatten().flatten() {
        file_dict.insert(file_name, url);
    }

    if file_dict.len() != identities.len() {
        eprintln!(
            "Warning: Some download URLs could not be retrieved. Expected {} but got {}.",
            identities.len(),
            file_dict.len()
        );
    }

    Ok(file_dict)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn test_gen_uwp_url_with_real_client() -> anyhow::Result<()> {
        let client = Client::builder()
            .danger_accept_invalid_certs(true)
            .build()?;

        // Test data
        let fulfillment_data = r#"{"ProductId":"9NBLGGH2JHXJ","WuBundleId":"96c73e7a-d14a-4aaa-9b59-5501d650b418","WuCategoryId":"d25480ca-36aa-46e6-b76b-39608d49558c","PackageFamilyName":"Microsoft.MinecraftUWP_8wekyb3d8bbwe","SkuId":"0011","Content":null,"PackageFeatures":null}"#;

        let result = gen_uwp_url(&client, fulfillment_data).await;

        match result {
            Ok(file_dict) => {
                assert!(
                    !file_dict.is_empty(),
                    "Expected non-empty file dictionary from real request."
                );
                println!("Successfully retrieved URLs: {:?}", file_dict);
            }
            Err(e) => {
                eprintln!("Test failed with error: {:?}", e);
                return Err(e);
            }
        }

        Ok(())
    }
}

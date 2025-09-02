use crate::genurl::Url;
use reqwest::Client;
use serde_json::Value;

/// Generates non UWP download URLs using the Microsoft Store delivery API.
///
/// # Arguments
/// * `client` - A reqwest client instance to make HTTP requests
/// * `product_id` - Product ID of the app
///
/// # Returns
/// * `Ok(HashMap<String, String>)` - Map of filenames to download URLs
/// * `Err(anyhow::Error)` - If any step in the process fails
///
pub async fn gen_non_uwp_url(client: &Client, product_id: &str) -> anyhow::Result<Url> {
    let resp = client
        .get(format!("https://storeedgefd.dsx.mp.microsoft.com/v9.0/packageManifests/{product_id}?market=US&locale=en-us&deviceFamily=Windows.Desktop"))
        .send()
        .await?;

    let data: Value = resp.json().await?;
    let data = data
        .get("Data")
        .and_then(|data| data.get("Versions"))
        .and_then(|versions| versions.as_array());

    let Some(data) = data else {
        anyhow::bail!("Invalid product ID");
    };

    let urls: Url = data
        .iter()
        .flat_map(|version| {
            let package_version = version
                .get("PackageVersion")
                .and_then(|str| str.as_str())
                .unwrap()
                .to_string();
            let package_name = version
                .get("DefaultLocale")
                .and_then(|default_locale| default_locale.get("PackageName"))
                .and_then(|str| str.as_str())
                .unwrap()
                .to_string();

            version
                .get("Installers")
                .and_then(|installers| installers.as_array())
                .unwrap()
                .iter()
                .map(move |installer| {
                    let installer_type = installer
                        .get("InstallerType")
                        .and_then(|str| str.as_str())
                        .unwrap()
                        .to_string();
                    let architecture = installer
                        .get("Architecture")
                        .and_then(|str| str.as_str())
                        .unwrap()
                        .to_string();
                    let url = installer
                        .get("InstallerUrl")
                        .and_then(|str| str.as_str())
                        .unwrap()
                        .to_string();

                    (
                        format!("{package_name}_{package_version}_{architecture}.{installer_type}"),
                        url,
                    )
                })
        })
        .collect();

    Ok(urls)
}

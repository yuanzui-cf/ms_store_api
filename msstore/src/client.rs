use anyhow::Result;
use reqwest::Client;
use serde_json::Value;

use crate::genurl::{gen_non_uwp_url, gen_uwp_url, Url};

/// Extracts fulfillment data from a JSON value representing product details.
///
/// # Arguments
/// * `value` - A JSON value containing product details
///
/// # Returns
/// An optional string containing the fulfillment data if found
fn get_fulfillment_data(value: &Value) -> Option<String> {
    let fulfillment_data = value
        .get("Payload")
        .and_then(|payload| payload.get("Skus"))
        .and_then(|skus| skus.get(0))
        .and_then(|sku| sku.get("FulfillmentData"))
        .and_then(|data| data.as_str());

    fulfillment_data.map(|str| str.to_string())
}

/// Fetches product details from Microsoft Store API and returns a URL for UWP app.
///
/// # Arguments
/// * `product_id` - The Microsoft Store product ID to look up
///
/// # Returns
/// A Result with the UWP app URL if successful
///
/// # Example
/// #[tokio::main]
/// async fn main() {
///     let product_id = "9NBLGGH4NNS1";
///     let url = fetch_product_details(product_id).await?;
/// }
/// ```
pub async fn fetch_product_details(product_id: &str) -> Result<Url> {
    let client = Client::builder()
        .danger_accept_invalid_certs(true)
        .build()?;

    let api_url = format!(
        "https://storeedgefd.dsx.mp.microsoft.com/v9.0/products/{}",
        product_id
    );

    let response = client
        .get(api_url)
        .query(&[
            ("market", "US"),
            ("locale", "en-us"),
            ("deviceFamily", "Windows.Desktop"),
        ])
        .header("Accept", "application/json")
        .send()
        .await?;

    let json_value: Value = response.json().await?;

    // 解析响应数据
    let fulfillment_data = get_fulfillment_data(&json_value);

    if let Some(fulfillment) = fulfillment_data
        && !fulfillment.is_empty()
    {
        return gen_uwp_url(&client, &fulfillment).await;
    }

    // 非UWP应用或无效产品ID
    gen_non_uwp_url(&client, product_id).await
}

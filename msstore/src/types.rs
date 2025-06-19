use serde::Deserialize;
use serde_json::Value;

#[derive(Debug, Deserialize)]
#[serde(rename_all = "PascalCase")]
pub struct FulfillmentDataContent {
    pub product_id: String,
    pub wu_bundle_id: String,
    pub wu_category_id: String,
    pub package_family_name: String,
    pub sku_id: String,
    pub content: Option<Value>,
    pub package_features: Option<Value>,
}

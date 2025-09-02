use std::sync::Arc;

use axum::http::{header, StatusCode};
use axum::response::IntoResponse;
use axum::{extract::Path, Extension, Json};
use msstore::client::fetch_product_details;
use serde::{Deserialize, Serialize};

use crate::conf;

#[derive(Deserialize, Serialize, Debug)]
pub struct Response {
    name: String,
    message: String,
    #[serde(skip_serializing_if = "serde_json::Value::is_null")]
    data: serde_json::Value,
}

pub async fn handler(
    Extension(config): Extension<Arc<conf::AppConfig>>,
    Path(id): Path<String>,
) -> impl IntoResponse {
    let mut res = Response {
        name: config.app_name.clone(),
        message: "".to_string(),
        data: serde_json::Value::Null,
    };

    let details = fetch_product_details(&id).await;

    if let Ok(data) = details {
        res.message = "Success".to_string();
        res.data = serde_json::to_value(data).unwrap();
    } else {
        res.message = format!("Error: {}", details.err().unwrap());
    }

    (
        StatusCode::OK,
        [(header::CONTENT_TYPE, "application/json; charset=utf-8")],
        Json(res),
    )
        .into_response()
}

pub async fn default_handler(Extension(config): Extension<Arc<conf::AppConfig>>) -> Json<Response> {
    Json(Response {
        name: config.app_name.clone(),
        message: "404 Not Found".to_string(),
        data: serde_json::Value::Null,
    })
}

use serde::Deserialize;

fn default_port() -> u16 {
    9000
}

fn default_app_name() -> String {
    "Macrohard Store API".to_string()
}

#[derive(Debug, Deserialize)]
pub struct AppConfig {
    #[serde(default = "default_app_name")]
    pub app_name: String,
    #[serde(default = "default_port")]
    pub port: u16,
}

use std::net::IpAddr;

use serde::Deserialize;

fn default_port() -> u16 {
    9000
}

fn default_app_name() -> String {
    "Macrohard Store API".to_string()
}

fn default_address() -> String {
    "0.0.0.0".to_string()
}

#[derive(Debug, Clone, Deserialize)]
pub struct AppConfig {
    #[serde(default = "default_app_name")]
    pub app_name: String,
    #[serde(default = "default_address")]
    pub address: String,
    #[serde(default = "default_port")]
    pub port: u16,
}

impl AppConfig {
    pub fn check(self) -> Self {
        AppConfig {
            app_name: self.app_name,
            address: if self.address.parse::<IpAddr>().is_ok() {
                self.address
            } else {
                "0.0.0.0".into()
            },
            port: if self.port > 0 {
                self.port
            } else {
                default_port()
            },
        }
    }
}

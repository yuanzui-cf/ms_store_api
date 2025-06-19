mod conf;

use anyhow::Ok;
use config::{Config, Environment, File};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let config = Config::builder()
        .add_source(File::with_name("config").required(false))
        .add_source(Environment::with_prefix("MSAPI").separator("_"))
        .build()?;

    let config = config.try_deserialize::<conf::AppConfig>()?;

    println!("{:?}", config);

    Ok(())
}

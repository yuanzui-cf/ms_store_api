pub mod conf;
mod handler;

use anyhow::Ok;

use axum::{Extension, Router};
use std::sync::Arc;
use tokio::signal;
use tracing::{info, warn};

use config::{Config, Environment, File};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    // Set logger
    let subscriber = tracing_subscriber::fmt().json().finish();
    tracing::subscriber::set_global_default(subscriber)?;

    let config = Config::builder()
        .add_source(File::with_name("config").required(false))
        .add_source(Environment::with_prefix("MSAPI").separator("_"))
        .build()?;

    let config = config.try_deserialize::<conf::AppConfig>()?;
    let config = Arc::new(config.check());

    let app = Router::new()
        .route("/{id}", axum::routing::any(handler::handler))
        .fallback(handler::default_handler)
        .layer(Extension(config.clone()));

    let listener =
        tokio::net::TcpListener::bind(format!("{}:{}", &config.address, config.port)).await?;

    info!("Listening on {}:{}", &config.address, config.port);

    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown_signal())
        .await?;

    Ok(())
}

async fn shutdown_signal() {
    let ctrl_c = async {
        signal::ctrl_c()
            .await
            .expect("Failed to listen for Ctrl+C signal.");
    };

    #[cfg(unix)]
    let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("Failed to listen for SIGTERM signal.")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {
            warn!("Strating shutting down...");
        },
        _ = terminate => {
            warn!("Strating shutting down...");
        },
    }
}

use std::collections::HashMap;

use rust_embed::Embed;

#[derive(Embed)]
#[folder = "assets/"]
pub struct Assets;

pub type Url = HashMap<String, String>;

package genurl

import "embed"

// Embed the data directory containing XML templates and other assets
//
//go:embed data
var DataDir embed.FS

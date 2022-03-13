package boilerplate

import "embed"

//go:embed config/*
var Config embed.FS

//go:embed .build/routes.json
var Routes []byte

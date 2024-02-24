//go:build !wasm && !js
// +build !wasm,!js

package main

import "embed"

//go:embed web
var embeddedWeb embed.FS

func init() {
	web = embeddedWeb
}

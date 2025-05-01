//go:build js
// +build js

package main

import (
	"github.com/avazquezcode/govetryx/internal/wasm"
)

func main() {
	// Register the functions that will be called from JavaScript
	wasm.RegisterFunctions()

	// Keep the program running
	select {}
} 
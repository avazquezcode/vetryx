//go:build js
// +build js

package wasm

import (
	"syscall/js"

	"github.com/avazquezcode/govetryx/internal/adapter/interpreter"
)

// CompileAndRun compiles and runs the given code, returning the output as a string.
func CompileAndRun(code string) (string, error) {
	return interpreter.RunCode(code)
}

// RegisterFunctions registers the Go functions to be called from JavaScript.
func RegisterFunctions() {
	js.Global().Set("compileAndRun", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return js.ValueOf(map[string]interface{}{
				"error": "Expected exactly one argument (code string)",
			})
		}

		code := args[0].String()
		output, err := CompileAndRun(code)
		if err != nil {
			return js.ValueOf(map[string]interface{}{
				"error": err.Error(),
			})
		}

		return js.ValueOf(map[string]interface{}{
			"output": output,
		})
	}))
} 
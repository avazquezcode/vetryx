package main

import (
	"log"
	"os"

	"github.com/avazquezcode/govetryx/internal/adapter/interpreter"
)

func main() {
	err := interpreter.RunFile(os.Args[1], os.Stdout)
	if err != nil {
		log.Fatalf("failed interpreting the script: %s", err.Error())
	}
}

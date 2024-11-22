package interpreter

import (
	"bytes"
	"fmt"
	"io"
	"os"

	interpreterpkg "github.com/avazquezcode/govetryx/internal/usecase/interpreter"
	"github.com/avazquezcode/govetryx/internal/usecase/parser"
	"github.com/avazquezcode/govetryx/internal/usecase/scanner"
)

// runCode triggers the interpreter to run the code.
func runCode(code []rune, stdout io.Writer) error {
	interpreter := interpreterpkg.NewInterpreter(stdout)
	s := scanner.NewScanner(code)
	tokens, err := s.Scan()
	if err != nil {
		return fmt.Errorf("failed on the lexer layer: %w", err)
	}

	p := parser.NewParser(tokens)
	statements, err := p.Parse()
	if err != nil {
		return err
	}

	resolver := interpreterpkg.NewResolver(interpreter)
	err = resolver.Resolve(statements)
	if err != nil {
		return fmt.Errorf("failed resolving the statements: %w", err)
	}

	return interpreter.Interpret(statements)
}

func RunFile(path string, stdout io.Writer) error {
	code, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed when reading the file: %w", err)
	}
	return runCode(bytes.Runes(code), stdout)
}

func RunCode(code string) (string, error) {
	var stdout bytes.Buffer
	err := runCode(bytes.Runes([]byte(code)), &stdout)
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}

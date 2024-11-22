package parser

import "fmt"

// ParsingErr represents an error during parsing.
type ParsingErr struct {
	errs []error
}

// NewParsingErr is a constructor for a parsing error.
func NewParsingErr(errs []error) *ParsingErr {
	return &ParsingErr{
		errs: errs,
	}
}

func (p *ParsingErr) Error() string {
	var message string

	for i, err := range p.errs {
		message = message + fmt.Sprintf("error #%d: %s \n", i+1, err.Error())
	}

	return message
}

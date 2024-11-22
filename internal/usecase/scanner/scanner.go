package scanner

import (
	"fmt"
	"strconv"

	"github.com/avazquezcode/govetryx/internal/domain/token"
)

// Scanner represents the scanner of the interpreter (aka: lexer).
type Scanner struct {
	sourceCode []rune
	tokens     []*token.Token
	start      int // represents the position where we start scanning a token.
	current    int // indicates the "pointer" position that moves forward during a token scan.
	line       int // indicates the line where we are standing on, during the scanning.
}

// NewScanner is a constructor for a Scanner.
func NewScanner(sourceCode []rune) *Scanner {
	return &Scanner{
		sourceCode: sourceCode,
		line:       1, // starts at line 1
	}
}

// Scan is the main method of the scanner.
// It scans all the tokens from a given sourceCode code.
func (s *Scanner) Scan() ([]*token.Token, error) {
	for !s.isEnd() {
		s.start = s.current
		if err := s.scanToken(); err != nil {
			return nil, fmt.Errorf("failed on lexer layer, while scanning line %d, with error: %w", s.line, err)
		}
	}

	// Add EOF
	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", nil, s.line))
	return s.tokens, nil
}

// scanToken scans the current token.
func (s *Scanner) scanToken() error {
	char := s.consume()

	if _, ok := singleChars[char]; ok {
		s.scanSingleChar(char)
		return nil
	}

	if _, ok := ignorableChars[char]; ok {
		return nil
	}

	if _, ok := matchableChars[char]; ok {
		s.scanMatchableChars(char)
		return nil
	}

	if char == '\n' {
		s.scanNewLine()
		return nil
	}

	if char == '"' {
		err := s.scanString()
		if err != nil {
			return fmt.Errorf("failed to scan string with err: %w", err)
		}
		return nil
	}

	if isDigit(char) {
		err := s.scanNumber()
		if err != nil {
			return fmt.Errorf("failed to scan number with err: %w", err)
		}

		return nil
	}

	if isLetter(char) {
		s.scanIdentifier()
		return nil
	}

	return fmt.Errorf("unexpected char")
}

// scanSingleChar handle the scanning of single chars (simple logic).
func (s *Scanner) scanSingleChar(char rune) {
	switch char {
	case '{':
		s.addToken(token.LeftBrace, nil)
	case '}':
		s.addToken(token.RightBrace, nil)
	case '(':
		s.addToken(token.LeftParentheses, nil)
	case ')':
		s.addToken(token.RightParentheses, nil)
	case ',':
		s.addToken(token.Comma, nil)
	case '+':
		s.addToken(token.Plus, nil)
	case '-':
		s.addToken(token.Minus, nil)
	case '*':
		s.addToken(token.Star, nil)
	case '/':
		s.addToken(token.Slash, nil)
	case '%':
		s.addToken(token.Modulus, nil)
	case '!':
		s.addToken(token.Bang, nil)
	case '#':
		// comments special handling
		for s.peek() != '\n' && !s.isEnd() {
			s.increment() // skip everything until the comment ends
		}
	case ';':
		s.addToken(token.Semicolon, nil)
	}
}

// scanMatchableChars handle the scanning of chars that can be matched with successive chars to form a composite token.
func (s *Scanner) scanMatchableChars(char rune) {
	switch char {
	case '<':
		if s.is('=') {
			s.increment()
			s.addToken(token.LowerOrEqual, nil)
			return
		}
		if s.is('>') {
			s.increment()
			s.addToken(token.NotEqual, nil)
			return
		}
		s.addToken(token.Lower, nil)
	case '>':
		if s.is('=') {
			s.increment()
			s.addToken(token.GreaterOrEqual, nil)
			return
		}
		s.addToken(token.Greater, nil)
	case '=':
		if s.is('=') {
			s.increment()
			s.addToken(token.EqualEqual, nil)
			return
		}
		s.addToken(token.Equal, nil)
	case '|':
		if s.is('|') {
			s.increment()
			s.addToken(token.Or, nil)
			return
		}
	case '&':
		if s.is('&') {
			s.increment()
			s.addToken(token.And, nil)
			return
		}
	case ':':
		if s.is('=') {
			s.increment()
			s.addToken(token.VarShortDeclarator, nil)
			return
		}
	}
}

// scanNewLine handles the scan of new lines.
func (s *Scanner) scanNewLine() {
	s.line++
	for s.peek() == '\n' {
		s.line++
		s.increment()
	}
}

// addToken creates a new token, and adds it to the tokens list.
func (s *Scanner) addToken(tokenType token.Type, literal interface{}) {
	lexeme := string(s.sourceCode[s.start:s.current])
	token := token.NewToken(tokenType, lexeme, literal, s.line)
	s.tokens = append(s.tokens, token)
}

// scanString handles the scanning of a string.
func (s *Scanner) scanString() error {
	for s.peek() != '"' && !s.isEnd() {
		if s.peek() == '\n' {
			// Multi line strings.
			s.line++
		}
		s.increment()
	}

	if s.isEnd() {
		return fmt.Errorf("missing quotes to close the string")
	}

	s.increment() // close the string (quotes)
	s.addToken(token.String, s.substring(s.start+1, s.current-1))
	return nil
}

// scanNumber handle the scanning of a number.
func (s *Scanner) scanNumber() error {
	for isDigit(s.peek()) {
		s.increment()
	}

	// check if is valid float
	if s.peek() == '.' && !isDigit(s.peekNext()) {
		return fmt.Errorf("the number is invalid")
	}

	// process float
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.increment() // read the "."
		for isDigit(s.peek()) {
			s.increment() // keep incrementing until the number ends
		}
	}

	value := s.substring(s.start, s.current)
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float with error: %w", err)
	}

	s.addToken(token.Number, number)
	return nil
}

// scanIdentifier handles the scanning of an identifier.
func (s *Scanner) scanIdentifier() {
	for isAlphaNum(s.peek()) {
		s.increment()
	}

	word := s.substring(s.start, s.current)
	tokenType, isReservedWord := token.ReservedWordsMapper[string(word)]
	if !isReservedWord {
		tokenType = token.Identifier // if not a reserved word, then is an identifier
	}
	s.addToken(tokenType, nil)
}

// isEnd determines if the scanner reached the end of the source code.
func (s *Scanner) isEnd() bool {
	return s.current >= len(s.sourceCode)
}

// peek returns the next char (rune) without moving the current pointer.
func (s *Scanner) peek() rune {
	if s.isEnd() {
		return 0
	}

	return s.sourceCode[s.current]
}

// consume consumes the next char (rune), by incrementing the pointer position, and returning the previous pointer position value.
func (s *Scanner) consume() rune {
	s.increment()
	return s.previous()
}

// previous returns the previous scanned char in the source code.
func (s *Scanner) previous() rune {
	return s.sourceCode[s.current-1]
}

// increment moves the current pointer of the scanner forward.
func (s *Scanner) increment() {
	s.current++
}

// is returns true if the current char (rune) matches with an expected char.
func (s *Scanner) is(expected rune) bool {
	if s.isEnd() {
		return false
	}

	if s.peek() != expected {
		return false
	}

	return true
}

// peekNext returns the char (rune) after the current one.
func (s *Scanner) peekNext() rune {
	nextIdx := s.current + 1
	if nextIdx >= len(s.sourceCode) {
		return 0
	}
	return s.sourceCode[nextIdx]
}

func (s *Scanner) previousToken() *token.Token {
	return s.tokens[len(s.tokens)-1]
}

func (s *Scanner) substring(start int, end int) string {
	return string(s.sourceCode[start:end])
}

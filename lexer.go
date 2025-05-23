package main

import (
	"strconv"
	"strings"
)

type Lexer struct {
	input    string
	position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) NextToken() Token {
	if l.position >= len(l.input) {
		return Token{Type: EOF, Value: ""}
	}
	l.skipWhitespace()
	currentChar := l.input[l.position]

	switch currentChar {
	case '{':
		l.position++
		return Token{Type: LEFT_BRACE, Value: "{"}
	case '}':
		l.position++
		return Token{Type: RIGHT_BRACE, Value: "}"}
	case '[':
		l.position++
		return Token{Type: LEFT_BRACKET, Value: "["}
	case ']':

		l.position++
		return Token{Type: RIGHT_BRACKET, Value: "]"}
	case ':':
		l.position++
		return Token{Type: COLON, Value: ":"}
	case ',':
		l.position++
		return Token{Type: COMMA, Value: ","}
	case '"':
		l.position++
		return l.readString()
	default:
		if isDigit(currentChar) {
			return l.readNumber()
		} else if strings.HasPrefix(l.input[l.position:], "true") {
			l.position += 4
			return Token{Type: BOOLEAN, Value: "true"}
		} else if strings.HasPrefix(l.input[l.position:], "false") {
			l.position += 5
			return Token{Type: BOOLEAN, Value: "false"}
		} else if strings.HasPrefix(l.input[l.position:], "null") {
			l.position += 4
			return Token{Type: NULL, Value: "null"}
		}
		return Token{Type: ILLEGAL, Value: string(currentChar)}
	}

}

func (l *Lexer) readString() Token {
	start := l.position
	for l.position < len(l.input) && l.input[l.position] != '"' {
		l.position++
	}
	if l.position >= len(l.input) {
		return Token{Type: ILLEGAL, Value: l.input[start:l.position]}
	}
	str := l.input[start:l.position]
	l.position++
	return Token{Type: STRING, Value: str}
}

func (l *Lexer) readNumber() Token {
	start := l.position
	for l.position < len(l.input) && (isDigit(l.input[l.position]) || l.input[l.position] == '.' || l.input[l.position] == '-') {
		l.position++
	}
	value := l.input[start:l.position]
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return Token{Type: ILLEGAL, Value: value}
	}
	return Token{Type: NUMBER, Value: value}
}

func (l *Lexer) skipWhitespace() {
	for l.position < len(l.input) && (l.input[l.position] == ' ' || l.input[l.position] == '\t' || l.input[l.position] == '\n' || l.input[l.position] == '\r') {
		l.position++
	}
}

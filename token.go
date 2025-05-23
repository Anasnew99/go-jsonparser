package main

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	STRING  TokenType = "STRING"
	NUMBER  TokenType = "NUMBER"
	BOOLEAN TokenType = "BOOLEAN"
	NULL    TokenType = "NULL"

	COLON TokenType = "COLON"
	COMMA TokenType = "COMMA"

	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"

	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"
)

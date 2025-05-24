package main

import "strconv"

// NewParser creates a new Parser instance by first tokenizing the input string
// and then initializing the parser with those tokens.
func NewParser(input string) *Parser {
	lexer := NewLexer(input)
	tokens := []Token{}
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		tokens = append(tokens, token)
	}
	return &Parser{
		tokens:  tokens,
		current: 0,
		input:   input,
	}
}

// Parser represents a JSON parser that processes tokens into a structured format.
// It maintains the current position in the token stream and the original input.
type Parser struct {
	tokens  []Token
	current int
	input   string
}

// Parse is the main entry point for parsing JSON input.
// It returns the parsed value as an interface{} (any) and any error encountered.
// The function determines whether to parse an object or array based on the first token.
func (p *Parser) Parse() (any, error) {
	if len(p.tokens) == 0 {
		return nil, ErrUnexpectedEOF
	}

	if p.currentToken().Type == LEFT_BRACE {
		obj, err := p.parseObject()
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
	if p.currentToken().Type == LEFT_BRACKET {
		arr, err := p.parseArray()
		if err != nil {
			return nil, err
		}
		return arr, nil
	}

	return nil, ErrUnexpectedToken
}

// parseObject parses a JSON object into a map[string]any.
// It handles key-value pairs separated by commas and enclosed in curly braces.
// Returns an error if the object structure is invalid.
func (p *Parser) parseObject() (map[string]any, error) {
	output := make(map[string]any)

	if p.consume(LEFT_BRACE) {
		for !p.consume(RIGHT_BRACE) {
			keyToken, err := p.expect(STRING)
			if err != nil {
				return nil, ErrExpectedString
			}
			key := keyToken.Value
			if p.consume(COLON) {
				value, err := p.parseValue()
				if err != nil {
					return nil, ErrExpectedValue
				}
				output[key] = value
			}
			if !p.consume(COMMA) && p.currentToken().Type != RIGHT_BRACE {
				return nil, ErrExpectedComma
			}
		}
	}
	return output, nil
}

// parseArray parses a JSON array into a []any.
// It handles comma-separated values enclosed in square brackets.
// Returns an error if the array structure is invalid.
func (p *Parser) parseArray() ([]any, error) {
	output := make([]any, 0)

	if p.consume(LEFT_BRACKET) {
		for !p.consume(RIGHT_BRACKET) {
			value, err := p.parseValue()
			if err != nil {
				return nil, ErrExpectedValue
			}
			output = append(output, value)
			if !p.consume(COMMA) && p.currentToken().Type != RIGHT_BRACKET {
				return nil, ErrExpectedComma
			}
		}
	}
	return output, nil
}

// parseValue handles parsing of individual JSON values.
// It supports strings, numbers, booleans, null, objects, and arrays.
// Returns the parsed value as an interface{} and any error encountered.
func (p *Parser) parseValue() (any, error) {
	switch p.currentToken().Type {
	case STRING:
		p.consume(STRING)
		return p.prevToken().Value, nil
	case NUMBER:
		p.consume(NUMBER)
		return strconv.ParseFloat(p.prevToken().Value, 64)
	case BOOLEAN:
		p.consume(BOOLEAN)
		return strconv.ParseBool(p.prevToken().Value)
	case NULL:
		p.consume(NULL)
		return nil, nil
	case LEFT_BRACE:
		return p.parseObject()
	case LEFT_BRACKET:
		return p.parseArray()
	default:
		return nil, ErrUnexpectedToken
	}
}

// currentToken returns the token at the current position in the token stream.
func (p *Parser) currentToken() Token {
	return p.tokens[p.current]
}

// prevToken returns the token at the previous position in the token stream.
func (p *Parser) prevToken() Token {
	return p.tokens[p.current-1]
}

// consume advances the parser if the current token matches the expected type.
// Returns true if the token was consumed, false otherwise.
func (p *Parser) consume(t TokenType) bool {
	if p.currentToken().Type == t {
		p.current++
		return true
	}
	return false
}

// expect attempts to consume a token of the specified type.
// Returns the consumed token if successful, or an error if the token type doesn't match.
func (p *Parser) expect(tokenType TokenType) (Token, error) {
	if p.consume(tokenType) {
		return p.prevToken(), nil
	}
	return Token{}, ErrUnexpectedToken
}

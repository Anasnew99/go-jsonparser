# JSON Parser

A simple JSON lexer implementation in Go that tokenizes JSON input into its constituent parts.

## Overview

This project implements a lexical analyzer (lexer) for JSON data. It takes JSON input as a string and breaks it down into tokens, which are the smallest meaningful units in the JSON syntax. The lexer supports all basic JSON data types including:

- Strings
- Numbers (integers and floating-point)
- Booleans (true/false)
- Null values
- Objects (key-value pairs)
- Arrays

## Features

- Tokenizes JSON input into meaningful components
- Handles whitespace and special characters
- Supports all standard JSON data types
- Simple and easy to understand implementation

## Project Structure

- `main.go` - Contains the main program entry point and example usage
- `lexer.go` - Implements the JSON lexer functionality
- `token.go` - Defines token types and structures

## Usage

Here's a simple example of how to use the JSON lexer:

```go
lexer := NewLexer(`{"name": "John", "age": 30, "isStudent": false, "skills": ["JavaScript", "Python", "Go"]}`)
for {
    token := lexer.NextToken()
    if token.Type == EOF {
        break
    }
    fmt.Println(token)
}
```

## Token Types

The lexer recognizes the following token types:

- `STRING` - JSON string values
- `NUMBER` - Numeric values (integers and floating-point)
- `BOOLEAN` - true/false values
- `NULL` - null values
- `COLON` - Key-value separator (:)
- `COMMA` - Value separator (,)
- `LEFT_BRACE` - Object start ({)
- `RIGHT_BRACE` - Object end (})
- `LEFT_BRACKET` - Array start ([)
- `RIGHT_BRACKET` - Array end (])
- `EOF` - End of input
- `ILLEGAL` - Invalid tokens

## Requirements

- Go 1.x or later

## Building and Running

To build and run the project:

```bash
go build
./jsonparser
```

## License

This project is open source and available under the MIT License. 
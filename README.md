# JSON Parser

A simple and efficient JSON parser written in Go that can parse JSON strings and files into Go data structures.

## Features

- Parses JSON strings and files
- Supports all standard JSON data types:
  - Objects (maps)
  - Arrays
  - Strings
  - Numbers
  - Booleans
  - Null
- Pretty printing of parsed JSON
- Error handling for invalid JSON

## Usage

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    // Parse a JSON string
    jsonStr := `{"name": "John", "age": 30, "isStudent": false, "height": 1.75, "skills": ["JavaScript", "Python", "Go"]}`
    parser := NewParser(jsonStr)
    jsonObject, err := parser.Parse()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Access parsed data
    name := jsonObject.(map[string]any)["name"].(string)
    fmt.Println("Name:", name)

    // Pretty print the parsed JSON
    prettyJSON, err := json.MarshalIndent(jsonObject.(map[string]any), "", "  ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Pretty Print:", string(prettyJSON))

    // Parse a JSON file
    fileParser := NewParser(string(fileContent))
    fileObject, err := fileParser.Parse()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    // ... work with fileObject
}
```

## Implementation Details

The parser uses a recursive descent approach with the following components:

1. **Lexer**: Tokenizes the input string into JSON tokens
2. **Parser**: Converts tokens into Go data structures
   - `Parse()`: Main entry point for parsing
   - `parseObject()`: Handles JSON objects
   - `parseArray()`: Handles JSON arrays
   - `parseValue()`: Handles primitive values

## Error Handling

The parser provides clear error messages for common JSON parsing issues:
- Unexpected end of file
- Invalid token types
- Missing commas
- Invalid object/array structure

## Project Structure

- `main.go` - Contains the main program entry point and example usage
- `lexer.go` - Implements the JSON lexer functionality
- `token.go` - Defines token types and structures

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
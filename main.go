package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	lexer := NewLexer(`{"name": "John", "age": 30, "isStudent": false, "height": 1.75, "skills": ["JavaScript", "Python", "Go"]}`)
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		fmt.Println(token)
	}

	parser := NewParser(`{"name": "John", "age": 30, "isStudent": false, "height": 1.75, "skills": ["JavaScript", "Python", "Go"]}`)
	jsonObject, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonObject)
	fmt.Println(jsonObject.(map[string]any)["name"].(string))
	prettyJSON, err := json.MarshalIndent(jsonObject.(map[string]any), "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Pretty Print:", string(prettyJSON))

	// read file
	file, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileParser := NewParser(string(file))
	fileObject, err := fileParser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(fileObject)
	prettyJSON, err = json.MarshalIndent(fileObject.(map[string]any), "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Pretty Print:", string(prettyJSON))
}

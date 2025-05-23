package main

import "fmt"

func main() {
	lexer := NewLexer(`{"name": "John", "age": 30, "isStudent": false, "height": 1.75, "skills": ["JavaScript", "Python", "Go"]}`)
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		fmt.Println(token)
	}
}

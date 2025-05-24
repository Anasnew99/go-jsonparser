package main

import "errors"

var (
	ErrUnexpectedEOF   = errors.New("UNEXPECTED EOF")
	ErrUnexpectedToken = errors.New("UNEXPECTED TOKEN")
	ErrExpectedString  = errors.New("EXPECTED STRING")
	ErrExpectedValue   = errors.New("EXPECTED VALUE")
	ErrExpectedComma   = errors.New("EXPECTED COMMA")
	ErrExpectedColon   = errors.New("EXPECTED COLON")
)

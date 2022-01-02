package compiler

import (
	"errors"
)

const (
	// Math Operators
	T_PLUS = iota
	T_MINUS
	T_MULT
	T_DIV
	//Binary Operators
	T_AND
	T_OR
	T_IS //Equivilent to == in other languages
	T_GT

	//Keywords
	T_IF
	T_FOR
	T_IN

	//Other/General
	T_NEWLN
	T_PAREN_L
	T_PAREN_R
	T_EQ
)

type Token struct {
	Name  string
	Value string
	Code  int64
}

func Tokenize(code *string) []Token {
	tokens := []Token{}
	// for i, c := range *code {

	// }
	return tokens
}

func charLookup(code *string, index int) (rune, error) {
	runes := []rune(*code)
	if index < 0 || index > len(runes)-1 {
		return 0, errors.New("index out of bounds")
	}
	return runes[index], nil
}
func isWordMatch(str string, code *string, index int) bool {
	return false
}

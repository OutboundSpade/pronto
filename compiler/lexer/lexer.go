package lexer

import (
	"fmt"
	"log"
	"regexp"
)

const (
	// Math Operators
	T_PLUS = iota
	T_MINUS
	T_MULT
	T_DIV

	//Binary Operators
	T_AND //&&
	T_OR  //||
	T_IS  //Equivilent to == in other languages
	T_GT
	T_GT_EQ
	T_LT
	T_LT_EQ

	//Keywords
	T_IF
	T_FOR
	T_IN
	T_FROM

	//Number
	T_NUM
	T_IDENT

	//Other/General
	T_NEWLN
	T_PAREN_L
	T_PAREN_R
	T_EQ
	T_UNKNOWN
)

type Token struct {
	Code  int64
	Value string
	Name  string
}

func Tokenize(code *string) []Token {
	tokens := []Token{}
	runes := []rune(*code)
	runes = append(runes, '\n')
	for cCounter := 0; cCounter < len(runes); cCounter++ {
		switch runes[cCounter] {
		case ' ':
			continue
		case '\t':
			continue
		case '\r':
			continue
		case '\n':
			tokens = append(tokens, Token{T_NEWLN, "", `NEWLINE`})
			continue
		case '+':
			tokens = append(tokens, Token{T_PLUS, "", "PLUS"})
			continue
		case '-':
			tokens = append(tokens, Token{T_MINUS, "", "MINUS"})
			continue
		case '*':
			tokens = append(tokens, Token{T_MULT, "", "MULT"})
			continue
		// case '/':
		// 	tokens = append(tokens, Token{T_DIV, "", "DIV"})
		// 	continue
		case '(':
			tokens = append(tokens, Token{T_PAREN_L, "", "L_PAREN"})
			continue
		case ')':
			tokens = append(tokens, Token{T_PAREN_R, "", "R_PAREN"})
			continue
		case '>':
			tokens = append(tokens, Token{T_GT, "", "GREATER_TN"})
			continue
		case '<':
			tokens = append(tokens, Token{T_LT, "", "LESS_TN"})
			continue

		}
		t := getFullTerm(&runes, cCounter)
		if len(t) > 0 {
			// log.Printf("WORD %v", t)
			if isComment(t) {
				l, c := traceLocFromIndex(&runes, skipComment(&runes, cCounter))
				log.Printf(`comment: %v endloc: %d,%d`, t, l, c)
				cCounter = skipComment(&runes, cCounter)
				continue
			}
			if isNumber(t) {
				tokens = append(tokens, Token{T_NUM, t, "NUMBER"})
				// log.Printf("Num: '%v'", t)
				cCounter += len(t) - 1
				continue
			}
			if isValidIdent(t) {
				tokens = append(tokens, Token{T_IDENT, t, "IDENT"})
				// log.Printf("Ident: '%v'", t)
				cCounter += len(t) - 1
				continue
			}
		}
		if runes[cCounter] == '/' {
			tokens = append(tokens, Token{T_DIV, "", "DIV"})
			continue
		}

		if t == "" {
			t = string(runes[cCounter])
		}
		unexpectedToken(&Token{T_UNKNOWN, t, ""}, &runes, cCounter)
	}
	return tokens
}

func charLookup(code *[]rune, index int) (rune, error) {
	if index < 0 || index > len(*code)-1 {
		return 0, fmt.Errorf("index out of bounds at %d", index)
	}
	return (*code)[index], nil
}
func getFullTerm(code *[]rune, startingIndex int) string {
	var t []rune
	patt := regexp.MustCompile(`[^+\-/*(){}\s]`)
	// patt := regexp.MustCompile(`\S`)

	for i := startingIndex; ; i++ {
		nextChar, err := charLookup(code, i)
		// fmt.Printf("%d : %s\n", nextChar, string(nextChar))
		if err != nil {
			panic(err)
		}
		if nextChar == '/' {
			n, _ := charLookup(code, i+1)
			if n == '/' || n == '*' {
				return string(nextChar) + string(n)
			}
		}
		match := patt.MatchString(string(nextChar))
		if err != nil {
			return ""
		}
		if match {
			t = append(t, nextChar)
			continue
		}
		break
	}
	return string(t)
}

//returns the location (index) after the comment ends
func skipComment(code *[]rune, startingIndex int) int {
	secChar, err := charLookup(code, startingIndex+1)
	if err != nil {
		panic(fmt.Errorf("error wall skipping comment:%v", err))

	}
	//If the second character is star -> multiline comment
	if secChar == '*' {
		for i := startingIndex + 2; i < len(*code)-1; i++ {
			charA, _ := charLookup(code, i)
			charB, _ := charLookup(code, i+1)
			if charA == '*' && charB == '/' {
				return i + 1
			}
		}
	}

	//for single line comment (//)
	for i := startingIndex; ; i++ {
		cChar, err := charLookup(code, i)
		if err != nil {
			panic(fmt.Errorf("error wall skipping singleline comment:%v", err))
		}
		if cChar == '\n' {
			return i
		}
	}
}

/*
Takes:
	code - pointer to array of runes
	index - a 1 dimentional index of the code
Returns:
	line - line number (int)
	col - character number (int)
*/
func traceLocFromIndex(code *[]rune, index int) (line int, col int) {
	line++
	var lastlc int = 0 // index of last newline
	for i := 0; i <= index; i++ {
		if c, e := charLookup(code, i); c == '\n' {
			if e != nil {
				panic(e)
			}
			lastlc = i
			line++
		}
	}
	col = index - lastlc
	return
}

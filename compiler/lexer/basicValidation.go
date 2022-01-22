/*
Basic validation for simple identifiers and numbers

advanced validation is handled at the parsing stage
*/
package lexer

import (
	"fmt"
	"regexp"
)

var (
	numExp   = regexp.MustCompile(`^[\d]*.{0,1}[\d]+$`)
	identExp = regexp.MustCompile(`^[_a-zA-z][_a-zA-z0-9]*$`)
	// commentExp = regexp.MustCompile(`^(\/\/[\s\S]*)|(\/\*[\s\S]*\*\/)$`)

	keyTerms = []string{
		"if",
		"else",
		"and",
		"or",
		"not",
		"for",
		"of",
	}
)

func isNumber(t string) bool {
	return numExp.MatchString(t)
}
func isValidIdent(t string) bool {
	return identExp.MatchString(t)
}
func isComment(t string) bool {
	return t == "//" || t == "/*"
}
func isKeyTerm(t string) bool {
	for _, k := range keyTerms {
		if k == t {
			return true
		}
	}
	return false
}

func unexpectedToken(t *Token, code *[]rune, index int) {
	lc, cc := traceLocFromIndex(code, index)
	if t.Code == T_UNKNOWN {
		fmt.Println(fmt.Errorf("unknown token at line %d, column %d: '%s'", lc, cc, (*t).Value))
	} else {
		fmt.Println(fmt.Errorf("unexpected token '%s' at line %d, column %d: '%s'", (*t).Name, lc, cc, (*t).Value))
	}
	// os.Exit(1)
}

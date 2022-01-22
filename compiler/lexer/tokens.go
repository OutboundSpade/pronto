package lexer

const (
	// Math Operators
	T_PLUS  = iota // +
	T_MINUS        // -
	T_MULT         // *
	T_DIV          // /

	//Boolean Operators
	T_AND   // and (&&)
	T_OR    // or (||)
	T_IS    // is (==)
	T_GT    // >
	T_GT_EQ // >=
	T_LT    // <
	T_LT_EQ // <=

	//Keywords
	T_IF   // if
	T_FOR  // for
	T_IN   // in
	T_FROM // from

	//Literals
	T_STR_LIT
	T_NUM_LIT

	T_IDENT

	//Brackets/braces
	T_PAREN_L // (
	T_PAREN_R // )
	T_CURL_L  // {
	T_CURL_R  // }
	T_SQR_L   // [
	T_SQR_R   // ]

	//Other/General
	T_NEWLN // \n
	T_EQ    // =

	T_UNKNOWN
)

type Token struct {
	Code  int64
	Value string
	Name  string
}

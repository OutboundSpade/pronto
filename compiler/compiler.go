package compiler

import "github.com/OutboundSpade/pronto/compiler/lexer"

type Token lexer.Token

var Tokenize = lexer.Tokenize

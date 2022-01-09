package parser

import "github.com/OutboundSpade/pronto/compiler"

type AST struct {
}

type Node struct {
	Parent   *Node
	Children *[]Node
}

func Parse(tokens *[]compiler.Token) {

}

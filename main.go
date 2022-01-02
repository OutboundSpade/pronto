package main

import (
	"log"

	"github.com/OutboundSpade/pronto/compiler"
	fileloader "github.com/OutboundSpade/pronto/fileLoader"
)

func main() {
	var code string = string(*fileloader.Load("test.pronto"))
	log.SetFlags(log.Lshortfile)

	tokens := compiler.Tokenize(&code)
	for _, token := range tokens {
		log.Printf("%v: %v\n", token.Name, token.Value)
	}
}

package main

import (
	"log"

	"github.com/OutboundSpade/pronto/compiler"
	fileloader "github.com/OutboundSpade/pronto/loader"
)

func main() {
	var code string = string(*fileloader.Load("test.pronto"))
	log.SetFlags(log.Lshortfile)

	log.Println(code)
	tokens := compiler.Tokenize(&code)
	for _, token := range tokens {
		log.Printf("%v: %v", token.Name, token.Value)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/OutboundSpade/pronto/compiler"
	fileloader "github.com/OutboundSpade/pronto/loader"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fpath := flag.Arg(0)

	var code string = string(*fileloader.Load(fpath))
	log.SetFlags(log.Lshortfile)

	fmt.Println(code)
	tokens := compiler.Tokenize(&code)
	// fmt.Printf("tokens: %v\n", tokens)
	for _, token := range tokens {
		fmt.Printf("%v", token.Name, token.Value)
	}
}

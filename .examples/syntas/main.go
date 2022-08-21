package main

import (
	"fmt"
	"github.com/alecthomas/chroma/lexers"
)

func main() {
	s := `
	asd
	a
	sd
	as
	da
	qqqqqqqqqqqqqqqqq
	
	*/
}`
	//fluent := str.NewFluent()
	get := lexers.Get("main.go")
	//quick.Highlight(fluent, s, get, "terminal256", "monokai")

	_ = s

	config := get.Config()
	fmt.Println(config.Name)
	fmt.Println(config.Filenames)
	fmt.Println(config.AliasFilenames)
	fmt.Println(config.NotMultiline)

	tokenise, _ := get.Tokenise(nil, s)
	for _, token := range tokenise.Tokens() {
		fmt.Println(token)
	}
}

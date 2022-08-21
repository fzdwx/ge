package main

import (
	"fmt"
	"github.com/fzdwx/ge/internal/views"
	"strings"
)

func main() {

	filenames := "README.md"
	document, err := views.LoadDocument(filenames)
	if err != nil {
		panic(err)
	}

	s := document.String()
	fmt.Println(s)
	sSp := strings.Split(s, "\n")
	_ = sSp
	fmt.Println("==============")

	render := document.Render()
	split := strings.Split(render, "\n")
	_ = split
	fmt.Println(render)
}

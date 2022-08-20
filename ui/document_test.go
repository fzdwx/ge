package ui

import (
	"fmt"
	"testing"
)

func Test_Load(t *testing.T) {

	//filenames := "document.go"
	filenames := "C:\\Users\\98065\\IdeaProjects\\ge\\README.md"
	document, err := loadDocument(filenames)
	if err != nil {
		panic(err)
	}

	fmt.Println(document)
	//
	//fmt.Println(document.Height())
	//fmt.Println("row height", document.Height(), "val:", document.Row(document.Height()))
	//fmt.Print(document.Row(4))
}

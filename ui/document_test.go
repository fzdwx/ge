package ui

import (
	"fmt"
	"testing"
)

func Test_Load(t *testing.T) {
	document, err := loadDocument("document.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(document)
}

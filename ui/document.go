package ui

import (
	"errors"
	"github.com/fzdwx/x/str"
	"os"
	"unicode/utf8"
)

type Document struct {
	rows []Row
}

func (d *Document) String() string {
	fluent := str.NewFluent()

	for _, row := range d.rows {
		fluent.Str(row.String())
	}

	return fluent.String()
}

func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Load(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(bytes) <= 0 {
		return err
	}

	var rows []Row
	var row Row

	// todo 默认作为utf-8
	for i, w := 0, 0; i < len(bytes); i += w {
		r, width := utf8.DecodeRune(bytes[i:])
		if r == utf8.RuneError {
			return errors.New("could not decode rune")
		}

		if r == '\x1b' && len(row) > 1 || r == '\n' {
			// a new key sequence has started
			rows = append(rows, row)
			row = []rune{}
		}

		row = append(row, r)
		w = width
	}

	d.rows = rows

	return nil
}

// loadDocument todo 暂时只加载一个
func loadDocument(filenames ...string) (*Document, error) {
	document := NewDocument()

	if len(filenames) <= 0 {
		return document, nil
	}

	if err := document.Load(filenames[0]); err != nil {
		return nil, err
	}

	return document, nil
}

package views

import (
	"github.com/fzdwx/ge/internal/syntax"
	"os"
)

type Document struct {
	Rows   Rows
	syntax syntax.Syntax
}

func (d *Document) String() string {
	return d.Rows.String()
}

func NewDocument() *Document {
	return &Document{Rows: Rows{}, syntax: syntax.From("")}
}

func (d *Document) Render() string {
	return d.syntax.Highlight(d.String())
}

func (d *Document) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	d.syntax = syntax.From(filename)

	if len(data) <= 0 {
		return err
	}

	rows, err := NewRows(data)
	if err != nil {
		return err
	}

	d.Rows = rows
	return nil
}

// Height get document Rows len.
func (d *Document) Height() int {
	return d.Rows.Len()
}

// Row get row by index
func (d *Document) Row(i int) Row {
	return d.Rows.Row(i)
}

// InsertRune insert rune at specified row and column
func (d *Document) InsertRune(r rune, row int, col int) {
	if r == '\n' {
		d.SplitLine(row, col)
		return
	}

	d.Rows.InsertRune(r, row, col)
}

func (d *Document) SplitLine(row int, col int) {
	d.Rows.SplitLine(row, col)
}

// Length  Value returns the value of the text input.
func (d *Document) Length() int {
	return d.Rows.TotalSize()
}

// LoadDocument todo 暂时只加载一个
// document is never null.
func LoadDocument(filenames ...string) (*Document, error) {
	document := NewDocument()

	if len(filenames) <= 0 {
		return document, nil
	}

	if err := document.Load(filenames[0]); err != nil {
		return document, err
	}

	return document, nil
}

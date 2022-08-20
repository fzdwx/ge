package views

import (
	"github.com/fzdwx/x/str"
	"os"
)

type Document struct {
	rows Rows
}

func (d *Document) Render() string {
	fluent := str.NewFluent()

	switch d.Height() {
	case 0:
		return str.Empty
	case 1:
		return d.rows[0].String()
	}

	fluent.Str(d.rows[0].String())
	for _, s := range d.rows[1:] {
		fluent.NewLine()
		fluent.Str(s.String())
	}

	return fluent.String()
}

func NewDocument() *Document {
	return &Document{rows: Rows{}}
}

func (d *Document) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(data) <= 0 {
		return err
	}

	rows, err := NewRows(data)
	if err != nil {
		return err
	}

	d.rows = rows
	return nil
}

// Height get document rows len.
func (d *Document) Height() int {
	return d.rows.Len()
}

func (d *Document) Row(i int) Row {
	return d.rows.Row(i)
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

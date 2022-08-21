package views

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/fzdwx/x/str"
	rw "github.com/mattn/go-runewidth"
	"unicode/utf8"
)

type (
	Row []rune

	Rows []Row
)

func (rs Rows) String() string {
	fluent := str.NewFluent()

	switch rs.Len() {
	case 0:
		return str.Empty
	case 1:
		return rs[0].String()
	}

	fluent.Str(rs[0].String())
	for _, s := range rs[1:] {
		fluent.NewLine()
		fluent.Str(s.String())
	}

	return fluent.String()
}

func NewRows(data []byte) (Rows, error) {
	reader := bufio.NewReader(bytes.NewBuffer(data))

	var rows []Row
	var row Row
	for {
		b, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		for i, w := 0, 0; i < len(b); i += w {
			r, width := utf8.DecodeRune(b[i:])
			if r == utf8.RuneError {
				return rows, errors.New("could not decode rune")
			}

			row = append(row, r)
			w = width
		}

		if !prefix {
			rows = append(rows, row)
			row = []rune{}
		}
	}

	return rows, nil
}

func (r Row) String() string {
	return string(r)
}

func (r Row) Col(col int) rune {
	return r[col]
}

func (r Row) RuneWidth(col int) int {
	if len(r) <= col {
		return 0
	} // todo
	return rw.RuneWidth(r[col])
}

// Len get row len
func (rs Rows) Len() int {
	return len(rs)
}

// Row get row by idx
func (rs Rows) Row(idx int) Row {

	if idx > rs.Len()-1 {
		return Row{}
	}

	return rs[idx]
}

func (rs Rows) SplitLine(row int, col int) {
	rowLine := rs.Row(row)
	head, tailSrc := rowLine[:col], rowLine[col:]
	tail := make([]rune, len(tailSrc))
	copy(tail, tailSrc)

	_ = append(rs[:row+1], rs[row:]...)
	rs[row] = head
	rs[row+1] = tail
}

func (rs Rows) InsertRune(r rune, row int, col int) {
	rs[row] = append(rs[row][:col], append([]rune{r}, rs[row][col:]...)...)
}

// TotalSize get rune total width.
func (rs Rows) TotalSize() int {
	var l int
	for _, row := range rs {
		l += rw.StringWidth(string(row))
	}
	return l
}

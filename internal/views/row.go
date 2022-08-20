package views

import (
	"bufio"
	"bytes"
	"errors"
	"unicode/utf8"
)

type (
	Row []rune

	Rows []Row
)

func NewRows(data []byte) (Rows, error) {
	reader := bufio.NewReader(bytes.NewBuffer(data))

	var rows []Row
	var row Row
	for {
		b, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if !prefix {
			rows = append(rows, row)
			row = []rune{}
		}

		for i, w := 0, 0; i < len(b); i += w {
			r, width := utf8.DecodeRune(b[i:])
			if r == utf8.RuneError {
				return rows, errors.New("could not decode rune")
			}

			row = append(row, r)
			w = width
		}
	}
	rows = append(rows, row)

	return rows, nil
}

func (r Row) String() string {
	return string(r)
}

func (rs Rows) Len() int {
	return len(rs)
}

// Row get row by idx
func (rs Rows) Row(idx int) Row {

	if idx > rs.Len()-1 {
		return nil
	}

	return rs[idx]
}

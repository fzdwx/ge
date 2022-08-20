package ui

type (
	Row []rune
)

func (r Row) String() string {
	return string(r)
}

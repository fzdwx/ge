package syntax

import (
	"github.com/fzdwx/x/str"
	"path"
	"strings"
)

type (
	Syntax interface {
		Type() string

		Highlight(s string) string

		FileName() string
	}

	Creator func(filename string) Syntax
)

var (
	m = map[string]Creator{
		str.Empty: func(filename string) Syntax {
			return Default(filename)
		},
		".md": func(filename string) Syntax {
			return MarkerDown(filename)
		},
	}
)

func From(filename string) Syntax {
	ext := strings.ToLower(path.Ext(filename))

	if f, ok := m[ext]; ok {
		return f(filename)
	}

	return m[str.Empty](filename)
}

type Default string

func (d Default) FileName() string            { return string(d) }
func (d Default) Type() string                { return "unknown" }
func (d Default) Highlight(str string) string { return str }

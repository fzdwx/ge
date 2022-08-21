package syntax

import (
	"github.com/charmbracelet/glamour"
	str2 "github.com/fzdwx/x/str"
	"strings"
)

type MarkerDown string

func (m MarkerDown) FileName() string { return string(m) }
func (m MarkerDown) Type() string     { return "md" }
func (m MarkerDown) Highlight(str string) string {
	render, err := glamour.Render(str, "dark")
	if err != nil {
		return str
	}

	// markdown will add 3 more lines of blanks
	sp := strings.Split(render, str2.NewLine)

	return strings.Join(sp[1:len(sp)-2], str2.NewLine)
}

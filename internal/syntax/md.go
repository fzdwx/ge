package syntax

import "github.com/charmbracelet/glamour"

type MarkerDown string

func (m MarkerDown) FileName() string { return string(m) }
func (m MarkerDown) Type() string     { return "md" }
func (m MarkerDown) Highlight(str string) string {
	render, err := glamour.Render(str, "dark")
	if err != nil {
		return str
	}
	return render
}

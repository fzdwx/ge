package syntax

type MarkerDown string

func (m MarkerDown) FileName() string { return string(m) }
func (m MarkerDown) Type() string     { return "md" }
func (m MarkerDown) Highlight(s string) string {
	//render, err := glamour.Render(s, "dark")
	//if err != nil {
	//	return s
	//}
	//
	//// markdown will add 3 more lines of blanks
	//sp := strings.Split(render, str.NewLine)
	//
	//return strings.Join(sp[1:len(sp)-2], str.NewLine)
	return s
}

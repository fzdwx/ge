package teax

import tea "github.com/charmbracelet/bubbletea"

type (
	// ErrorMsg error msg
	ErrorMsg struct {
		Err error
	}
)

func Check(err error) tea.Cmd {
	return func() tea.Msg {
		if err != nil {
			return ErrorMsg{Err: err}
		}
		return nil
	}
}

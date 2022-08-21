package teax

import tea "github.com/charmbracelet/bubbletea"

type batch struct {
	cmds []tea.Cmd
}

func Batch(cmds ...tea.Cmd) *batch {
	return &batch{cmds: cmds}
}

func (b *batch) Append(cmd tea.Cmd) *batch {
	b.cmds = append(b.cmds, cmd)
	return b
}

func (b *batch) Cmd() tea.Cmd {
	return tea.Batch(b.cmds...)
}

func (b *batch) Check(err error) *batch {
	b.Append(Check(err))
	return b
}

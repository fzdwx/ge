package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Keymap struct {
	quit key.Binding
}

func NewKeymap() *Keymap {
	return &Keymap{
		quit: key.NewBinding(
			key.WithKeys(tea.KeyCtrlC.String()),
			key.WithHelp(tea.KeyCtrlC.String(), "quit program"),
		),
	}
}

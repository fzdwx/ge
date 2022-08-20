package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/ge/config"
)

type Ui struct {
	cfg      *config.Config
	document *Document

	Program *tea.Program
	Keymap  *Keymap
}

func New(cfg *config.Config) *Ui {
	this := &Ui{Keymap: NewKeymap(), cfg: cfg}
	//document, err := loadDocument(cfg.Filenames)

	return this
}

func (u Ui) Init() tea.Cmd {
	return nil
}

func (u Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, u.Keymap.quit):
			return u, tea.Quit
		}
	}
	return u, nil
}

func (u Ui) View() string {
	return "hello world"
}

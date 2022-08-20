package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/ge/config"
	"github.com/fzdwx/ge/internal/teax"
	"github.com/fzdwx/ge/internal/views"
)

type Ui struct {
	cfg      *config.Config
	document *views.Document

	Program *tea.Program
	Keymap  *Keymap
}

func New(cfg *config.Config) *Ui {
	this := &Ui{Keymap: NewKeymap(), cfg: cfg}
	return this
}

func (u *Ui) Init() tea.Cmd {
	document, err := views.LoadDocument(u.cfg.Filenames...)
	u.document = document
	return teax.Check(err)
}

func (u *Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, u.Keymap.quit):
			return u, tea.Quit
		}
	case teax.ErrorMsg:
		// todo handle error msg
	}
	return u, nil
}

func (u *Ui) View() string {
	return u.document.Render()
}

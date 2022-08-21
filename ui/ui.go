package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/ge/config"
	"github.com/fzdwx/ge/internal/teax"
	"github.com/fzdwx/ge/internal/views"
	"github.com/fzdwx/x/str"
)

type (
	Ui struct {
		cfg      *config.Config
		document *views.Document

		Program *tea.Program
		Keymap  *Keymap

		termSize
	}

	termSize struct {
		Width  int
		Height int
	}
)

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
	case tea.WindowSizeMsg:
		u.termSize = termSize{Width: msg.Width, Height: msg.Height}
	case teax.ErrorMsg:
		// todo handle error msg
	}
	return u, nil
}

func (u *Ui) View() string {
	return u.document.Render() + str.NewLine + u.termSize.String()
}

func (t termSize) String() string {
	return fmt.Sprintf("w:%d - h:%d", t.Width, t.Height)
}

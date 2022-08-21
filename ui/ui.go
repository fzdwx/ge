package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/ge/config"
	"github.com/fzdwx/ge/internal/teax"
	"github.com/fzdwx/ge/internal/views"
)

type (
	Ui struct {
		cfg *config.Config

		// current document
		document *views.Document

		textarea *Textarea

		Program *tea.Program
		Keymap  *Keymap
	}
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238"))

	blurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.HiddenBorder())
)

func New(cfg *config.Config) *Ui {
	area := NewTextArea()
	area.ShowLineNumbers = true
	area.Cursor.Style = cursorStyle
	area.FocusedStyle.Base = focusedBorderStyle
	area.BlurredStyle.Base = blurredBorderStyle
	area.Focus()
	this := &Ui{
		Keymap:   NewKeymap(),
		textarea: area,
		cfg:      cfg,
	}
	return this
}

func (u *Ui) Init() tea.Cmd {
	batch := teax.Batch(Blink)

	document, err := views.LoadDocument(u.cfg.Filenames...)
	u.document = document
	u.textarea.SetDocument(u.document)
	batch.Check(err)
	return batch.Cmd()
}

func (u *Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	batch := teax.Batch()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, u.Keymap.quit):
			return u, tea.Quit
		}
	case tea.WindowSizeMsg:
		u.textarea.SetHeight(msg.Height - 2)
		u.textarea.SetWidth(msg.Width)
	case teax.ErrorMsg:
		// todo handle error msg
	}

	textarea, cmd := u.textarea.Update(msg)
	batch.Append(cmd)
	u.textarea = textarea

	return u, batch.Cmd()
}

func (u *Ui) View() string {
	return u.textarea.View()
}

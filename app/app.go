package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/ge/config"
	"github.com/fzdwx/ge/ui"
)

type App struct {
	ui *ui.Ui
}

func New(filenames []string) *App {
	return &App{ui: ui.New(config.New(filenames))}
}

func (a App) StartUp(ops ...tea.ProgramOption) error {
	a.ui.Program = tea.NewProgram(a.ui, ops...)
	return a.ui.Program.Start()
}

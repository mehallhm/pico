package engine

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Quit(_ *Dataframe, _ []string) (EngineModel, error) {
	model := QuitModel("exiting...")

	return model, nil
}

type QuitModel string

func (m QuitModel) Init() tea.Cmd {
	return tea.Quit
}

func (m QuitModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	return m, tea.Quit
}

func (m QuitModel) View() string {
	return string(m)
}

func (m QuitModel) Focus() {}

func (m QuitModel) Blur() {}

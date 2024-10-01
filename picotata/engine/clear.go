package engine

import tea "github.com/charmbracelet/bubbletea"

func Clear(_ *Dataframe, _ []string) (EngineModel, error) {
	return &clearModel{}, nil
}

type clearModel struct{}

func (m *clearModel) Init() tea.Cmd {
	return nil
}

func (m *clearModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	return m, nil
}

func (m *clearModel) View() string {
	return ""
}

func (m *clearModel) Focus() {}

func (m *clearModel) Blur() {}

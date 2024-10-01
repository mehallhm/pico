package engine

import tea "github.com/charmbracelet/bubbletea"

// TODO: This REALLY should handle strings... at least parse them or something
type Dataframe struct {
	Columns []string
	Data    map[string][]float64
}

type EngineModel interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (EngineModel, tea.Cmd)
	View() string
	Focus()
	Blur()
}

type TextModel struct {
	Text string
}

func (m TextModel) Init() tea.Cmd {
	return nil
}

func (m TextModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	return m, nil
}

func (m TextModel) View() string {
	return m.Text
}

func (m TextModel) Focus() {}

func (m TextModel) Blur() {}

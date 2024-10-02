package engine

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Count(df *Dataframe, _ []string) (EngineModel, error) {
	if len(df.Columns) == 0 {
		return nil, fmt.Errorf("data not loaded")
	}

	model := CountModel{
		text:    fmt.Sprintf("Count: %v", len(df.Data[df.Columns[0]])),
		focused: false,
	}
	return model, nil
}

type CountModel struct {
	text    string
	focused bool
}

func (m CountModel) Init() tea.Cmd {
	return nil
}

func (m CountModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	return m, nil
}

func (m CountModel) View() string {
	return m.text
}

func (m CountModel) Focus() {}

func (m CountModel) Blur() {}

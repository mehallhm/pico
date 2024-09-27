package engine

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Browse(df *Dataframe) (*BrowseModel, error) {

	columns := []table.Column{}

	for _, name := range df.Columns {
		columns = append(columns, table.Column{Title: name, Width: 10})
	}

	rows := []table.Row{}
	for rowIdx := 0; rowIdx < len(df.Data[columns[0].Title]); rowIdx++ {
		row := make(table.Row, 0, len(df.Data[columns[0].Title]))
		for _, col := range df.Data {
			row = append(row, fmt.Sprintf("%v", col[rowIdx]))
		}

		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(false),
		table.WithRows(rows),
		table.WithHeight(20),
	)
	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("200")).
		Bold(false)
	t.SetStyles(s)

	return &BrowseModel{
		table: t,
	}, nil
}

type BrowseModel struct {
	table table.Model
}

func (m *BrowseModel) Init() tea.Cmd {
	return nil
}

func (m *BrowseModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *BrowseModel) View() string {
	return m.table.View() + "\n" + m.table.HelpView()
}

func (m *BrowseModel) Focus() {
	m.table.Focus()
}

func (m *BrowseModel) Blur() {
	m.table.Blur()
}

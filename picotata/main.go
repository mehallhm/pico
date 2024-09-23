package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mehallhm/picotata/engine"
)

type errMsg error

func main() {
	p := tea.NewProgram(initalModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type model struct {
	textInput textinput.Model
	output    string
	table     table.Model
	showTable bool
	err       error
}

func initalModel() model {
	ti := textinput.New()
	ti.Placeholder = "type a command or `quit` to exit"
	ti.Focus()
	ti.Prompt = "| "
	ti.CharLimit = 256
	ti.Width = 50

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case cmdMsg:
		switch msg.form {
		case engine.TextDisplay:
			m.table.Blur()
			m.showTable = false
			m.textInput.Focus()
			m.output = msg.msg
		case engine.TableDisplay:
			columns := []table.Column{}
			for _, name := range msg.data.Columns {
				columns = append(columns, table.Column{Title: name, Width: 10})
			}

			rows := []table.Row{}
			for rowIdx := 0; rowIdx < len(msg.data.Data[columns[0].Title]); rowIdx++ {
				row := make(table.Row, 0, len(msg.data.Data[columns[0].Title]))
				for _, col := range msg.data.Data {
					row = append(row, fmt.Sprintf("%v", col[rowIdx]))
				}

				rows = append(rows, row)
			}

			t := table.New(
				table.WithColumns(columns),
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
				Foreground(lipgloss.Color("240")).
				Background(lipgloss.Color("0")).
				Bold(false)
			t.SetStyles(s)

			m.showTable = true
			m.table = t

		}
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyTab:
			if m.showTable {
				if m.table.Focused() {
					m.table.Blur()
					m.textInput.Focus()

					s := table.DefaultStyles()

					s.Header = s.Header.
						BorderStyle(lipgloss.NormalBorder()).
						BorderForeground(lipgloss.Color("240")).
						BorderBottom(true).
						Bold(false)
					s.Selected = s.Selected.
						Foreground(lipgloss.Color("229")).
						Bold(false)

					m.table.SetStyles(s)
				} else {
					m.textInput.Blur()
					m.table.Focus()

					s := table.DefaultStyles()

					s.Header = s.Header.
						BorderStyle(lipgloss.NormalBorder()).
						BorderForeground(lipgloss.Color("240")).
						BorderBottom(true).
						Bold(false)
					s.Selected = s.Selected.
						Foreground(lipgloss.Color("229")).
						Background(lipgloss.Color("57")).
						Bold(false)

					m.table.SetStyles(s)
				}
			}

			return m, textinput.Blink

		case tea.KeyEnter:
			val := m.textInput.Value()
			m.textInput.SetValue("")
			return m, executeCmd(val)
		}

		if m.showTable {
			m.table, cmd = m.table.Update(msg)
			cmds = append(cmds, cmd)
		}

		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil

	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	if m.showTable {
		return fmt.Sprintf(
			"Input a command\n\n%s\n\n%s\n",
			m.textInput.View(),
			m.output,
		) + baseStyle.Render(m.table.View()) + "\n"
	}

	return fmt.Sprintf(
		"Input a command\n\n%s\n\n%s\n",
		m.textInput.View(),
		m.output,
	) + "\n"
}

type cmdMsg struct {
	msg  string
	data engine.Dataframe
	form engine.DisplayType
}

func executeCmd(cmd string) tea.Cmd {
	return func() tea.Msg {
		statement, err := engine.PrepareStatement(cmd)
		if err != nil {
			return errMsg(err)
		}

		out, err := engine.ExecuteStatement(statement)
		if err != nil {
			return errMsg(err)
		}

		return cmdMsg{
			msg:  out.Text,
			form: out.Form,
			data: out.Data,
		}
	}
}

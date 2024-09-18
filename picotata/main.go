package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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
	err       error
}

func initalModel() model {
	ti := textinput.New()
	ti.Placeholder = "help"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

type cmdMsg string

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

		return cmdMsg(out)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case cmdMsg:
		m.output = string(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			val := m.textInput.Value()
			m.textInput.SetValue("")
			return m, executeCmd(val)
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Input a command\n\n%s\n\nOut:\n %s\n",
		m.textInput.View(),
		m.output,
	) + "\n"
}

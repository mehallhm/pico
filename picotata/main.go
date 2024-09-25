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

	outputModel engine.EngineModel

	err error
}

func initalModel() model {
	ti := textinput.New()
	ti.Placeholder = "type a command or `quit` to exit"
	ti.Focus()
	ti.Prompt = "| "
	ti.CharLimit = 256
	ti.Width = 50

	return model{
		textInput:   ti,
		err:         nil,
		outputModel: engine.TextModel{Text: "  •         \n┏┓┓┏┏┓╋┏┓╋┏┓\n┣┛┗┗┗┛┗┗┻┗┗┻\n┛           \na tiny stata clone\n\n"},
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
		m.outputModel = msg
		cmd = m.outputModel.Init()
		return m, cmd

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyTab:
			if m.textInput.Focused() {
				m.textInput.Blur()
				m.outputModel.Focus()

				return m, nil
			} else {
				m.textInput.Focus()
				m.outputModel.Blur()

				return m, textinput.Blink
			}

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

	m.outputModel, cmd = m.outputModel.Update(msg)
	cmds = append(cmds, cmd)

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"Input a command\n\n%s\n\n%s\n",
		m.textInput.View(),
		m.outputModel.View(),
	)
}

type cmdMsg engine.EngineModel

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

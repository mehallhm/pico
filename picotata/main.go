package main

import (
	"fmt"
	"os"
	"strings"

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

	outputModel engine.EngineModel

	err error
}

var promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("200")).Bold(true)

var logoStyle = lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("61"))

func initalModel() model {
	ti := textinput.New()
	ti.Placeholder = "type a command or `quit` to exit"
	ti.Focus()
	ti.Prompt = promptStyle.Render(">> ")
	ti.CharLimit = 255
	ti.Width = 50

	return model{
		textInput:   ti,
		err:         nil,
		outputModel: engine.TextModel{Text: logoStyle.Render("  •         \n┏┓┓┏┏┓╋┏┓╋┏┓\n┣┛┗┗┗┛┗┗┻┗┗┻\n┛           \na tiny stata clone\n\n")},
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
			if m.textInput.Focused() {
				val := m.textInput.Value()
				m.textInput.SetValue("")
				return m, executeCmd(val)
			}
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

var headerTextStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))
var headerSlashesStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

func (m model) View() string {
	header := headerTextStyle.Render("Picotata ") + headerSlashesStyle.Render(strings.Repeat("/", 71))

	// text := ""
	// for i := range 16 {
	// 	a := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(fmt.Sprintf("%v", i)))
	// 	text = text + a.Render(fmt.Sprintf("///// Text - %v /////\n", i))
	// }
	return header + "\n" + fmt.Sprintf(
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

package engine

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

func Load(df *Dataframe, args []string) (*LoadModel, error) {
	if len(args) != 1 {
		fp := filepicker.New()
		fp.AllowedTypes = []string{".csv"}
		fp.Height = 10
		fp.CurrentDirectory, _ = os.UserHomeDir()

		return &LoadModel{fp: fp}, nil
	}

	f, err := os.Open(args[0])
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	df.Columns = records[0]
	table := make(map[string][]float64, len(records[0]))

	for colIdx := 0; colIdx < len(records[0]); colIdx++ {
		for rowIdx := 1; rowIdx < len(records); rowIdx++ {
			val, err := strconv.ParseFloat(strings.TrimSpace(records[rowIdx][colIdx]), 64)
			if err != nil {
				continue
			}
			table[records[0][colIdx]] = append(table[records[0][colIdx]], val)
		}
	}

	df.Data = table

	return &LoadModel{text: "Loaded!", df: df}, nil
}

type LoadModel struct {
	df           *Dataframe
	text         string
	fp           filepicker.Model
	selectedFile string
	focus        bool

	err error
}

func (m *LoadModel) Init() tea.Cmd {
	return m.fp.Init()
}

func (m *LoadModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
		if !m.focus {
			return m, nil
		}
	}

	m.fp, cmd = m.fp.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.fp.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
		m, m.err = Load(df, []string{path})
		return m, nil
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.fp.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, cmd
	}

	return m, cmd
}

func (m *LoadModel) View() string {
	var s strings.Builder
	s.WriteString("\n  ")
	s.WriteString(m.text)
	if m.fp.Height > 0 {
		if m.err != nil {
			s.WriteString(m.fp.Styles.DisabledFile.Render(m.err.Error()))
		} else if m.selectedFile == "" {
			s.WriteString("Pick a file:")
		} else {
			s.WriteString("Selected file: " + m.fp.Styles.Selected.Render(m.selectedFile))
		}
		s.WriteString("\n" + m.fp.View())
	}
	return s.String()
}

func (m *LoadModel) Focus() {
	m.focus = true
}

func (m *LoadModel) Blur() {
	m.focus = false
}

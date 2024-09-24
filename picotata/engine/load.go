package engine

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func Load(df *Dataframe, args []string) (LoadModel, error) {
	if len(args) != 1 {
		return LoadModel{}, fmt.Errorf("incorrect arguments")
	}

	f, err := os.Open(args[0])
	if err != nil {
		return LoadModel{}, err
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return LoadModel{}, err
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

	return LoadModel{text: "Loaded!", focused: false}, nil
}

type LoadModel struct {
	text    string
	focused bool
}

func (m LoadModel) Init() tea.Cmd {
	return nil
}

func (m LoadModel) Update(msg tea.Msg) (EngineModel, tea.Cmd) {
	return m, nil
}

func (m LoadModel) View() string {
	return m.text
}

func (m LoadModel) Focus() {}

func (m LoadModel) Blur() {}

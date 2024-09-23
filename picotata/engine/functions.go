package engine

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Load(args []string) (*Dataframe, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("incorrect arguments")
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

	df := Dataframe{
		Columns: records[0],
	}

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

	return &df, nil
}

func Count(df *Dataframe) (int, error) {
	return len(df.Data[df.Columns[0]]), nil
}

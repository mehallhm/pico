package engine

type StatementType int

const (
	InfoStatement = iota
	HelpStatement
	LoadStatement
	CountStatement
)

type Statement struct {
	Type StatementType
	Args []string
}

type Dataframe struct {
	Columns []string
	Data    map[string][]float64
}

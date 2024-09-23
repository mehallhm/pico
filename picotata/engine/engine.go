package engine

type StatementType int

const (
	QuitStatement = iota
	InfoStatement
	HelpStatement
	LoadStatement
	CountStatement
	SummarizeStatement
)

type Statement struct {
	Type StatementType
	Args []string
}

// TODO: This REALLY should handle strings... at least parse them or something
type Dataframe struct {
	Columns []string
	Data    map[string][]float64
}

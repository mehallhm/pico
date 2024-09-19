package engine

import (
	"fmt"
)

var df *Dataframe = nil

type DisplayType int

const (
	TextDisplay = iota
	FileDisplay
	TableDisplay
)

type FunctionReturn struct {
	Text string
	Form DisplayType
	Data Dataframe
}

func ExecuteStatement(statement *Statement) (FunctionReturn, error) {
	switch statement.Type {
	case InfoStatement:
		return FunctionReturn{Form: TextDisplay, Text: "  •         \n┏┓┓┏┏┓╋┏┓╋┏┓\n┣┛┗┗┗┛┗┗┻┗┗┻\n┛           \na tiny stata clone\n\n"}, nil
	case HelpStatement:
		return FunctionReturn{Form: TextDisplay, Text: "no help for you"}, nil
	case LoadStatement:
		df, _ = Load(statement.Args)
		return FunctionReturn{Form: TableDisplay, Data: *df}, nil
	case CountStatement:
		count, err := Count(df)
		return FunctionReturn{Form: TextDisplay, Text: fmt.Sprintf("%v", count)}, err
	}

	return FunctionReturn{Form: TextDisplay, Text: "cmd not found"}, nil
}

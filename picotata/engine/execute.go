package engine

import (
	"fmt"
)

var df *Dataframe = nil

func ExecuteStatement(statement *Statement) (string, error) {
	switch statement.Type {
	case InfoStatement:
		return "  •         \n┏┓┓┏┏┓╋┏┓╋┏┓\n┣┛┗┗┗┛┗┗┻┗┗┻\n┛           \na tiny stata clone\n\n", nil
	case HelpStatement:
		return "no help for you", nil
	case LoadStatement:
		df, _ = Load(statement.Args)
		return fmt.Sprint(df), nil
	case CountStatement:
		count, err := Count(df)
		return fmt.Sprintf("Count: %v", count), err
	}

	return "command not found", nil
}

package engine

var df *Dataframe = &Dataframe{}

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

func ExecuteStatement(statement *Statement) (EngineModel, error) {
	switch statement.Type {
	case QuitStatement:
		return Quit()
	case ClearStatement:
		return Clear()
	// case HelpStatement:
	// 	return FunctionReturn{Form: TextDisplay, Text: "no help for you"}, nil
	case LoadStatement:
		return Load(df, statement.Args)
	case CountStatement:
		return Count(df)
	case BrowseStatement:
		return Browse(df)
		// case SummarizeStatement:
		// 	return FunctionReturn{Form: TextDisplay, Text: "counts here"}, nil
	default:
		return TextModel{Text: "cmd not found"}, nil
	}
}

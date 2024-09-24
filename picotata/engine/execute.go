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
	// case InfoStatement:
	// 	return FunctionReturn{Form: TextDisplay, Text: "  •         \n┏┓┓┏┏┓╋┏┓╋┏┓\n┣┛┗┗┗┛┗┗┻┗┗┻\n┛           \na tiny stata clone\n\n"}, nil
	// case HelpStatement:
	// 	return FunctionReturn{Form: TextDisplay, Text: "no help for you"}, nil
	case LoadStatement:
		return Load(df, statement.Args)
	case CountStatement:
		return Count(df)
		// case SummarizeStatement:
		// 	return FunctionReturn{Form: TextDisplay, Text: "counts here"}, nil
	}

	// return FunctionReturn{Form: TextDisplay, Text: "cmd not found"}, nil
	return CountModel{text: "cmd not found", focused: false}, nil
}

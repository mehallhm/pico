package engine

import (
	"strings"
)

func PrepareStatement(input string) (*Statement, error) {
	input = strings.TrimSpace(input)
	switch {
	case strings.HasPrefix(input, "quit"):
		return &Statement{
			Type: QuitStatement,
		}, nil
	case strings.HasPrefix(input, "info"):
		return &Statement{
			Type: InfoStatement,
		}, nil
	case strings.HasPrefix(input, "help"):
		return &Statement{
			Type: HelpStatement,
		}, nil
	case strings.HasPrefix(input, "load"):
		args := strings.Split(input, " ")
		return &Statement{
			Type: LoadStatement,
			Args: args[1:],
		}, nil
	case strings.HasPrefix(input, "count"):
		return &Statement{
			Type: CountStatement,
		}, nil
	case strings.HasPrefix(input, "summarize"):
		return &Statement{
			Type: SummarizeStatement,
		}, nil
	case strings.HasPrefix(input, "browse"):
		return &Statement{
			Type: BrowseStatement,
		}, nil
	default:
		return &Statement{
			Type: MissingStatement,
		}, nil
	}
}

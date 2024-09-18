package engine

import (
	"fmt"
	"strings"
)

func PrepareStatement(input string) (*Statement, error) {
	input = strings.TrimSpace(input)
	switch {
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
	default:
		return nil, fmt.Errorf("unknown command")
	}
}

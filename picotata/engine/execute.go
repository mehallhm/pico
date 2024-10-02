package engine

import (
	"fmt"
	"strings"
)

var df *Dataframe = &Dataframe{}

type Command struct {
	command string
	def     func(df *Dataframe, args []string) (EngineModel, error)
}

var Cmds []Command = []Command{
	{"quit", Quit},
	{"clear", Clear},
	{"load", Load},
	{"browse", Browse},
	{"count", Count},
}

func ExecuteV(statement string) (EngineModel, error) {
	input := strings.TrimSpace(statement)
	args := strings.Split(input, " ")
	for _, cmd := range Cmds {
		if strings.HasPrefix(input, cmd.command) {
			return cmd.def(df, args[1:])
		}
	}

	return nil, fmt.Errorf("cmd not found")
}

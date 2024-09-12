package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/mehallhm/picosql/compiler"
)

func main() {
	Start()
}

func Start() {
	fmt.Println("repl started. type `.quit` to exit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			slog.Error("error reading repl input")
			return
		}

		// meta commands
		if input[0] == '.' {
			doMetaCommand(input)
			continue
		}

		statement, err := compiler.PrepareStatement(input)
		if err != nil {
			fmt.Println("invalid command")
			continue
		}

		compiler.ExecuteStatement(statement)
	}

}

func doMetaCommand(command string) {
	switch command {
	case ".quit\n":
		os.Exit(0)
	case ".help\n":
		fmt.Println("no help for you")
	case ".info\n":
		fmt.Print("  •      ┓\n┏┓┓┏┏┓┏┏┓┃\n┣┛┗┗┗┛┛┗┫┗\n┛       ┗\na tiny sqlite clone\n\n")
	default:
		fmt.Println("command not found")
	}
}

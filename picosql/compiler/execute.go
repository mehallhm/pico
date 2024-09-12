package compiler

import (
	"fmt"
)

func ExecuteStatement(statement *Statement) {
	switch statement.StatementType {
	case SelectStatement:
		fmt.Println("do a select")
	case InsertStatement:
		fmt.Println("do an insert")
	}
}

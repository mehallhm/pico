package compiler

import (
	"fmt"
	"strconv"
	"strings"
)

type StatementType int

const (
	SelectStatement = iota
	InsertStatement
)

type InsertionRow struct {
	Id       uint32
	Username [32]byte
	Email    [255]byte
}

type Statement struct {
	StatementType StatementType
	Row           InsertionRow
}

func PrepareStatement(statement string) (*Statement, error) {
	switch {
	case strings.HasPrefix(statement, "SELECT"):
		statement = strings.ReplaceAll(statement, "\n", "")
		args := strings.Split(statement, " ")

		if len(args) < 4 {
			return nil, fmt.Errorf("invalid arguments")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, fmt.Errorf("invalid arguments")
		}

		var username [32]byte
		copy(username[:], args[2])

		var email [255]byte
		copy(email[:], args[3])

		return &Statement{
			StatementType: SelectStatement,
			Row: InsertionRow{
				Id:       uint32(id),
				Username: username,
				Email:    email,
			},
		}, nil
	case strings.HasPrefix(statement, "INSERT"):
		return &Statement{
			StatementType: InsertStatement,
		}, nil
	default:
		return nil, fmt.Errorf("syntax error")
	}
}

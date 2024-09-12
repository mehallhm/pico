package compiler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	in  string
	out StatementType
}

var tests = []test{{"select", 0}, {"insert", 1}}

func TestPrepareStatement(t *testing.T) {
	_, err := PrepareStatement("bing bong")
	assert.Equal(t, err, fmt.Errorf("syntax error"))

	for _, test := range tests {
		statement, err := PrepareStatement(test.in)
		assert.Nil(t, err, "error preparing statement")

		assert.Equal(t, statement.StatementType, test.out, "statement preperation fail")
	}
}

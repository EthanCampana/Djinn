package ast

import (
	"djinn/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&CreateStatement{
				Token: token.Token{Type: token.CREATE, Literal: "cr"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "cr myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}

}

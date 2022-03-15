package lexer

import (
	"djinn/token"
	"testing"
)

//TODO ADD MORE TOKENS AND Write up a TEST FOR DJINN Syntax
func TestNextToken(t *testing.T) {
	input := `cr input = 10;
	fn add_two(x){
		x = x + 2
		mu x
	};
	input = add_two(input)	
	"foobar"
	"foo bar"
	[1,2];
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CREATE, "cr"},
		{token.IDENT, "input"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "add_two"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.MUTE, "mu"},
		{token.IDENT, "x"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "input"},
		{token.ASSIGN, "="},
		{token.IDENT, "add_two"},
		{token.LPAREN, "("},
		{token.IDENT, "input"},
		{token.RPAREN, ")"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - TokenType wrong. expected=%q, got=%q, %q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

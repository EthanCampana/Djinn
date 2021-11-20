package colormanager

import (
	"djinn/lexer"
	"djinn/token"
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func TestGenerateColor(t *testing.T) {

	assertCorrectColor := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Fatalf("expected=%q, got=%q", want, got)
		}
	}

	t.Run("Testing Two Colors", func(t *testing.T) {
		line := "hello_world = yep"
		cyan := color.New(color.FgCyan).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()
		want := fmt.Sprintf("%s %s %s \n", white("hello_world"), cyan("="), white("yep"))
		l := lexer.New(line)
		cm := New()
		tok_l := BuildTokenArray(l)
		got := cm.GenerateColor(tok_l)
		assertCorrectColor(t, got, want)
	})

}

func BuildTokenArray(l *lexer.Lexer) []token.Token {
	token_slice := []token.Token{}
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		token_slice = append(token_slice, tok)
	}
	return token_slice

}

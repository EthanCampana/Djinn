package repl

import (
	"bufio"
	"djinn/colormanager"
	"djinn/lexer"
	"djinn/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	cm := colormanager.New()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		toks := []token.Token{}
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
			toks = append(toks, tok)
		}
		fmt.Fprint(out, cm.GenerateColor(toks))

	}
}

package repl

import (
	"bufio"
	"djinn/colormanager"
	"djinn/evaluator"
	"djinn/lexer"
	"djinn/parser"
	"fmt"
	"io"

	"github.com/fatih/color"
)

const PROMPT = ">> "

const whoops = `
 __          ___                          /\//\/|
 \ \        / / |                        |/\//\/ 
  \ \  /\  / /| |__   ___   ___  _ __  ___       
   \ \/  \/ / | '_ \ / _ \ / _ \| '_ \/ __|      
    \  /\  /  | | | | (_) | (_) | |_) \__ \      
     \/  \/   |_| |_|\___/ \___/| .__/|___/      
                                | |              
                                |_|              
`

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
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors(), cm)
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			cm.Print(out, evaluated.Inspect())
			cm.Print(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string, cm *colormanager.ColorManager) {
	// io.WriteString(out, whoops)
	cm.PrintC(out, whoops, color.New(color.FgHiYellow))
	for _, msg := range errors {
		cm.PrintC(out, "\t"+msg+"\n", color.New(color.FgHiRed))
	}

}

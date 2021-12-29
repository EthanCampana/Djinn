package repl

import (
	"bufio"
	"djinn/lexer"
	"djinn/parser"
	"fmt"
	"io"
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
	// cm := colormanager.New()

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
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, whoops)
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}

}

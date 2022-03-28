package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeygo/evaluator"
	"monkeygo/lexer"
	"monkeygo/object"
	"monkeygo/parser"
)

const MONKEY_FACE = `
  ___              ___  
 (o o)            (o o) 
(  V  ) monkeygo (  V  )
--m-m--------------m-m--
`
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		switch line {
		case "quit":
			io.WriteString(out, "bye")
			return
		}
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

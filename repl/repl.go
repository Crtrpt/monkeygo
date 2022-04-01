package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeygo/compiler"
	"monkeygo/lexer"
	"monkeygo/parser"
	"monkeygo/vm"
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
	// env := object.NewEnvironment()
	// macroEnv := object.NewEnvironment()
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
		//编译到bytecode
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}
		//执行bytecode
		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")

		// evaluator.DefineMacros(program, macroEnv)
		// expanded := evaluator.ExpandMacros(program, macroEnv)

		// evaluated := evaluator.Eval(expanded, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }
	}
}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"monkeygo/evaluator"
	"monkeygo/lexer"
	"monkeygo/object"
	"monkeygo/parser"
	"monkeygo/repl"
	"os"
	"os/user"
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
	macroEnv := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		switch line {
		case "quit":
			io.WriteString(out, "\n\nヾ(￣▽￣)后会有期")
			return
		case "exit":
			io.WriteString(out, "\n\nヾ(￣▽￣)后会有期")
			return
		}
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, " o(╥﹏╥)o 解析错误:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("你好,%s 欢迎使用mongkeygo编程语言!\n",
		user.Username)
	fmt.Printf("当前环境为script执行环境 版本:v0.0.1\n")
	repl.Start(os.Stdin, os.Stdout)
}

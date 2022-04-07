package compiler

import (
	"fmt"
	"monkeygo/ast"
	"monkeygo/code"
	"monkeygo/lexer"
	"monkeygo/object"
	"monkeygo/parser"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1; 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpPop),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpPop),
			}},
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
				code.Make(code.OpPop),
			},
		},

		{
			input:             "1 - 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpSub),
				code.Make(code.OpPop)},
		},
		{
			input:             "1 * 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpMul),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "2 / 1",
			expectedConstants: []interface{}{2, 1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpDiv),
				code.Make(code.OpPop),
			},
		},

		{
			input:             "-1",
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpMinus),
				code.Make(code.OpPop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()
	for _, tt := range tests {
		program := parse(tt.input)
		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("编译器错误: %s", err)
		}
		bytecode := compiler.Bytecode()
		//测试bytecode 和期望的是否一致
		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("指令错误: %s", err)
		}
		//测试结果
		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("结果错误: %s", err)
		}
	}
}

// 解析输入字符串
func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

//测试指令
func testInstructions(
	expected []code.Instructions,
	actual code.Instructions,
) error {
	concatted := concatInstructions(expected)
	if len(actual) != len(concatted) {
		return fmt.Errorf("指令长度错误.\n期待=\n%s\n实际 = \n%s",
			concatted, actual)
	}
	for i, ins := range concatted {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\n期待=\n%s\n实际 =\n%s",
				i, concatted, actual)
		}
	}
	return nil
}

//连接指令
func concatInstructions(s []code.Instructions) code.Instructions {
	out := code.Instructions{}
	for _, ins := range s {
		out = append(out, ins...)
	}
	return out
}

//测试常量
func testConstants(
	t *testing.T,
	expected []interface{},
	actual []object.Object,
) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d",
			len(actual), len(expected))
	}
	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed: %s",
					i, err)
			}
		}
	}
	return nil
}

//测试整数
func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
	return nil
}

func TestBooleanExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "true",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpTrue),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpFalse),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 > 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpGreaterThan),
				code.Make(code.OpPop),
			},
		},
		// {
		// 	input:             "1 < 2",
		// 	expectedConstants: []interface{}{2, 1},
		// 	expectedInstructions: []code.Instructions{
		// 		code.Make(code.OpConstant, 0),
		// 		code.Make(code.OpConstant, 1),
		// 		code.Make(code.OpGreaterThan),
		// 		code.Make(code.OpPop),
		// 	},
		// },
		// {
		// 	input:             "1 == 2",
		// 	expectedConstants: []interface{}{1, 2},
		// 	expectedInstructions: []code.Instructions{
		// 		code.Make(code.OpConstant, 0),
		// 		code.Make(code.OpConstant, 1),
		// 		code.Make(code.OpEqual),
		// 		code.Make(code.OpPop),
		// 	},
		// },
		// {

		// 	input:             "1 != 2",
		// 	expectedConstants: []interface{}{1, 2},
		// 	expectedInstructions: []code.Instructions{
		// 		code.Make(code.OpConstant, 0),
		// 		code.Make(code.OpConstant, 1),
		// 		code.Make(code.OpNotEqual),
		// 		code.Make(code.OpPop),
		// 	},
		// },
		// {
		// 	input:             "true == false",
		// 	expectedConstants: []interface{}{},
		// 	expectedInstructions: []code.Instructions{
		// 		code.Make(code.OpTrue),
		// 		code.Make(code.OpFalse),
		// 		code.Make(code.OpEqual),
		// 		code.Make(code.OpPop),
		// 	},
		// },
		// {
		// 	input:             "true != false",
		// 	expectedConstants: []interface{}{},
		// 	expectedInstructions: []code.Instructions{
		// 		code.Make(code.OpTrue),
		// 		code.Make(code.OpFalse),
		// 		code.Make(code.OpNotEqual),
		// 		code.Make(code.OpPop),
		// 	},
		// },
		// {
		// 	input:             "!true",
		// 	expectedConstants: []interface{}{},
		// 	expectedInstructions: []code.Instructions{
		// 		code.Make(code.OpTrue),
		// 		code.Make(code.OpBang),
		// 		code.Make(code.OpPop),
		// 	},
		// },
	}
	runCompilerTests(t, tests)
}

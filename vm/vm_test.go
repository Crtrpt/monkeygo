package vm

import (
	"fmt"
	"monkeygo/ast"
	"monkeygo/compiler"
	"monkeygo/lexer"
	"monkeygo/object"
	"monkeygo/parser"
	"testing"
)

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
	}
	return nil
}

type vmTestCase struct {
	input    string
	expected interface{}
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()
	for _, tt := range tests {
		//解析输入到 ast
		program := parse(tt.input)
		//初始化编译器
		comp := compiler.New()
		//编译器处理ast
		err := comp.Compile(program)
		//处理相应的结果
		if err != nil {
			t.Fatalf("编译器错误: %s", err)
		}
		//初始化虚拟机
		vm := New(comp.Bytecode())
		//虚拟机执行bytecode
		err = vm.Run()

		if err != nil {
			t.Fatalf("虚拟机错误: %s", err)
		}
		// stackElem := vm.StackTop()
		// testExpectedObject(t, tt.expected, stackElem)
		//弹出最后一个栈元素
		stackElem := vm.LastPoppedStackElem()
		//测试期待结果
		testExpectedObject(t, tt.expected, stackElem)
	}
}

func testExpectedObject(
	t *testing.T,
	expected interface{},
	actual object.Object,
) {
	t.Helper()
	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	case bool:
		err := testBooleanObject(bool(expected), actual)
		if err != nil {
			t.Errorf("testBooleanObject failed: %s", err)
		}
	}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3}, // FIXME
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
		{"50 / 2 * 2 + 10 - 5", 55},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"5 * (2 + 10)", 60},

		{"-5", -5},
		{"-10", -10},
		{"-50 + 100 + -50", 0},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	runVmTests(t, tests)
}

func TestBooleanExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		// {"1 > 2", false},
		// {"1 < 1", false},
		// {"1 > 1", false},
		// {"1 == 1", true},
		// {"1 != 1", false},
		// {"1 == 2", false},
		// {"1 != 2", true},
		// {"true == true", true},
		// {"false == false", true},
		// {"true == false", false},
		// {"true != false", true},
		// {"false != true", true},
		// {"(1 < 2) == true", true},
		// {"(1 < 2) == false", false},
		// {"(1 > 2) == true", false},
		// {"(1 > 2) == false", true},

		// {"!true", false},
		// {"!false", true},
		// {"!5", false},
		// {"!!true", true},
		// {"!!false", false},
		// {"!!5", true},
	}
	runVmTests(t, tests)
}

func testBooleanObject(expected bool, actual object.Object) error {
	result, ok := actual.(*object.Boolean)
	if !ok {
		return fmt.Errorf("object is not Boolean. got=%T (%+v)",
			actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
	}
	return nil
}

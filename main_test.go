package main

import (
	"fmt"
	"github.com/alex1sz/rpn_cli_calculator/calculator"
	"strconv"
	"testing"
)

var dispatchTable = calculator.DispatchTable{
	"+": func(token string, stack *calculator.Stack) {
		stack.Push(stack.PopToFloat() + stack.PopToFloat())
	},
	"-": func(token string, stack *calculator.Stack) {
		value := stack.PopToFloat()
		stack.Push(stack.PopToFloat() - value)
	},
	"*": func(token string, stack *calculator.Stack) {
		stack.Push(stack.PopToFloat() * stack.PopToFloat())
	},
	"/": func(token string, stack *calculator.Stack) {
		value := stack.PopToFloat()
		stack.Push(stack.PopToFloat() / value)
	},
	"FLOAT": func(token string, stack *calculator.Stack) {
		value, _ := strconv.ParseFloat(token, 64)
		stack.Push(value)
	},
	"__DEFAULT__": func(token string, stack *calculator.Stack) {
		fmt.Printf("Unknown token: %s \n enter a valid operator, number, or type 'help' for more info...\n", token)
	},
}

func TestCalcToShell(t *testing.T) {
	expression := []string{"4", "3", "ee", "+"}
	stack := new(calculator.Stack)
	// expect calcToShell to handle bad input mid slice and not panic
	calcToShell(&expression, dispatchTable, stack)
}

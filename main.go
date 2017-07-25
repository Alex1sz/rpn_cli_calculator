package main

import (
	"bufio"
	"fmt"
	"github.com/alex1sz/rpn_cli_calculator/calculator"
	"os"
	"strconv"
	"strings"
)

func calcToShell(expression *[]string, dispatchTable calculator.DispatchTable, stack *calculator.Stack) {
	calcVal := calculator.Evaluate(expression, dispatchTable, stack)
	fmt.Printf("%f \n", calculator.ToFloat(calcVal))
}

func main() {
	dispatchTable := calculator.DispatchTable{
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
		"q": func(token string, stack *calculator.Stack) {
			os.Exit(0)
		},
		"help": func(token string, stack *calculator.Stack) {
			fmt.Printf("'q' to exit \n operators implemented '+', '-', '*', '/'")
		},
	}
	// program name is always first argument
	fmt.Printf("Welcome to: %s\n ", os.Args[0])

	var expression []string
	stack := new(calculator.Stack)

	// check for os.Args and evaluate each
	if len(os.Args) > 1 {
		fmt.Printf("calc %s", strings.Join(os.Args[1:], " "))
		for _, argStr := range os.Args[1:] {

			expression = append(expression, argStr)
			calcToShell(&expression, dispatchTable, stack)
		}
	}

	// NewReader returns a new Reader whose buffer has the default size
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		// ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter.

		inputStr, _ := reader.ReadString('\n')
		// Replace() returns new copy of string with trailing new line removed
		inputStr = strings.Replace(inputStr, "\n", "", -1)
		// handle > one float/int per line
		parsedInputArr := strings.Split(inputStr, " ")
		expression = append(expression, parsedInputArr...)

		calcToShell(&expression, dispatchTable, stack)
	}
}

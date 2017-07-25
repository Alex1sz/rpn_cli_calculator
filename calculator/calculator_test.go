package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// helper method returns stack w/ element
func initStackWithElement() (stack *Stack) {
	stack = new(Stack)
	stack.Push("+")
	return
}

func TestPush(t *testing.T) {
	stack := initStackWithElement()
	stackLength := stack.Length()

	if stackLength != 1 {
		t.Errorf("Expected length to eq 1, Got: %d", stackLength)
	}
}

func TestPop(t *testing.T) {
	stack := initStackWithElement()
	popVal := toString(stack.Pop())

	if popVal != "+" {
		t.Errorf("Expected '+', Got: %s", popVal)
	}
}

func TestPopReturnsNilForEmptyStack(t *testing.T) {
	stack := new(Stack)
	value := stack.Pop()

	if &value == nil {
		t.Errorf("Expected nil, Got: %s", value)
	}
}

func TestToString(t *testing.T) {
	stack := new(Stack)
	// expect toString() to return empty string "" when called on any non string
	stack.Push(9)
	str := toString(stack)

	if str != "" {
		t.Errorf("Expected empty string, Got: %s", str)
	}
}

// test evaluate for single line expression
func TestEvaluate(t *testing.T) {
	dispatchTable := DispatchTable{
		"+": func(token string, stack *Stack) {
			stack.Push(stack.PopToFloat() + stack.PopToFloat())
		},
		"-": func(token string, stack *Stack) {
			value := stack.PopToFloat()
			stack.Push(stack.PopToFloat() - value)
		},
		"*": func(token string, stack *Stack) {
			stack.Push(stack.PopToFloat() * stack.PopToFloat())
		},
		"/": func(token string, stack *Stack) {
			value := stack.PopToFloat()
			stack.Push(stack.PopToFloat() / value)
		},
		"FLOAT": func(token string, stack *Stack) {
			value, _ := strconv.ParseFloat(token, 64)
			stack.Push(value)
		},
		"__DEFAULT__": func(token string, stack *Stack) {
			panic(fmt.Sprintf("Unkown token %q", token))
		},
	}

	expressionsMap := map[string]map[string]float64{
		"35,5": {
			"+": 40.0,
			"-": 30.0,
			"*": 175.0,
			"/": 7.0,
		},
		"9.0,2.75": {
			"+": 11.75,
			"-": 6.25,
			"*": 24.75,
			"/": 3.272727272727273,
		},
		"1": {
			"+": 1.0,
		},
		"5,8": {
			"+": 13.0,
			"-": -3.0,
		},
		"-2,-3": {
			"+": -5.0,
			"-": 1.0,
			"*": 6.0,
			"/": 0.6666666666666666,
		},
		"1,1,5": {
			"+": 6,
			"-": -4,
		},
	}

	for expressionStr, expectedResults := range expressionsMap {
		// loop thru 'expectedResults' by operator keys
		for operator, expected := range expectedResults {
			expression := strings.Split(expressionStr, ",")
			expression = append(expression, operator)

			result := Evaluate(&expression, dispatchTable, new(Stack))

			if result != expected {
				t.Errorf("Expected expression: '%s' to eq: %f, got: %f", strings.Join(expression, ","), expected, result)
			}
		}
	}
}

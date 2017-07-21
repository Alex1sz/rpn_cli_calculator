package calculator

import (
	"strconv"
)

type (
	DispatchFunc  func(string, *Stack)
	DispatchTable map[string]DispatchFunc
)

// evaluates takes the expression, branching logic (ie the dispatch table) & stack for state
func Evaluate(expression []string, dispatchTable DispatchTable, stack *Stack) interface{} {

	for _, token := range expression {
		var dispatchFunction DispatchFunc

		if _, err := strconv.ParseFloat(token, 64); err == nil {
			dispatchFunction = dispatchTable["FLOAT"]
		} else {
			var evalsOk bool
			if dispatchFunction, evalsOk = dispatchTable[token]; !evalsOk {
				dispatchFunction = dispatchTable["__DEFAULT__"]
			}
		}
		dispatchFunction(token, stack)
	}
	return stack.Pop()
}

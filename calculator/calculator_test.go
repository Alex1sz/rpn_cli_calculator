package calculator

import (
	"testing"
)

// helper method returns stack w/ element
func initStackWithElement() (stack *Stack) {
	stack = new(Stack)
	stack.Push("+")
	return
}

func TestStackPush(t *testing.T) {
	stack := initStackWithElement()
	stackLength := stack.Length()

	if stackLength != 1 {
		t.Errorf("Expected stack length to eq 1, Got: %d", stackLength)
	}
}

func TestStackPopWhenNotEmpty(t *testing.T) {
	stack := initStackWithElement()
	popVal := toString(stack.Pop())

	if popVal != "+" {
		t.Errorf("Expected '+', Got: %s", popVal)
	}
}

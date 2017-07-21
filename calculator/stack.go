package calculator

// LIFO stack backed by linked list
type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (stack *Stack) Length() int {
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &Element{value, stack.top}
	stack.size++
}

// Pop removes top element from stack & return its value or nil when empty
func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

// Godocs Interface conversions & type assertions https://golang.org/doc/effective_go.html#interface_conversions

// General typecasting functions
func toFloat(value interface{}) float64 {
	if float, ok := value.(float64); ok {
		return float
	}
	return 0.0
}

func toString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

func (s *Stack) PopToFloat() float64 {
	return toFloat(s.Pop())
}

// func toInterfaces(value interface{}) []interface{} {
// 	if v, ok := value.([]interface{}); ok {
// 		return v
// 	}
// 	return []interface{}{value}
// }

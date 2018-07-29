package main

import "fmt"

// IntStack is a structure that represent a stack made by slice
type IntStack struct {
	stack []int // null slice
}

// Push element into the stack
func (stack *IntStack) Push(item int) {
	stack.stack = append(stack.stack, item)
}

// Top return the top element without removeing it
func (stack *IntStack) Top() int {
	return stack.stack[len(stack.stack)-1]
}

// Pop remove the top element of the stack
func (stack *IntStack) Pop() {
	stack.stack = stack.stack[:len(stack.stack)-1]
}

//String return the stack represented as a string
func (stack *IntStack) String() string {
	return fmt.Sprintf("%v", stack.stack)
}

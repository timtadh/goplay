package main

import "stack"
import "fmt"

type IntStack stack.Stack

func (s *IntStack) Push(i int) {
	self := (*stack.Stack)(s)
	self.P__push(i)
}

func (s *IntStack) Pop() int {
	self := (*stack.Stack)(s)
	i, _ := self.P__pop().(int)
	return i
}

func main() {
	s := (*IntStack)(stack.NewStack())
	s.Push(1)
	fmt.Println(s.Pop())
}

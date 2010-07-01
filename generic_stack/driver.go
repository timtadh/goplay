package main

import "stack"
import "fmt"

type IntStack stack.Stack
type StrStack stack.Stack

func (s *IntStack) Peek() int {
	self := (*stack.Stack)(s)
	i, _ := self.Peek().(int)
	return i
}

func (s *IntStack) Push(i int) {
	self := (*stack.Stack)(s)
	self.Push(i)
}

func (s *IntStack) Pop() int {
	self := (*stack.Stack)(s)
	i, _ := self.Pop().(int)
	return i
}

func (s *StrStack) Peek() string {
	self := (*stack.Stack)(s)
	i, _ := self.Peek().(string)
	return i
}

func (s *StrStack) Push(i string) {
	self := (*stack.Stack)(s)
	self.Push(i)
}

func (s *StrStack) Pop() string {
	self := (*stack.Stack)(s)
	i, _ := self.Pop().(string)
	return i
}

func main() {
	{
		s := (*IntStack)(stack.NewStack())
		s.Push(1)
		fmt.Println(s.Peek())
		s.Push(2)
		fmt.Println(s.Pop())
		fmt.Println(s.Pop())
	}
	{
		s := (*StrStack)(stack.NewStack())
		s.Push("Hello")
		fmt.Println(s.Peek())
		s.Push("World")
		fmt.Println(s.Pop())
		fmt.Println(s.Pop())
	}
}

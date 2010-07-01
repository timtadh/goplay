package stack

import "container/list"

type Stack struct {
    list *list.List
}

func NewStack() *Stack {
    self := new(Stack)
    self.list = list.New()
    return self
}

func (self *Stack) Empty() bool { return self.list.Len() <= 0 }

func (self *Stack) P__peek() interface{} {
    e := self.list.Front()
    t, _ := e.Value.(interface{})
    return t
}

func (self *Stack) P__push(t interface{}) {
    self.list.PushFront(t)
}

func (self *Stack) P__pop() interface{} {
    e := self.list.Front()
    t, _ := e.Value.(interface{})
    self.list.Remove(e)
    return t
}

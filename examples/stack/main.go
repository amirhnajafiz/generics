// Implementing Stack with Golang generics
package main

import "fmt"

type Data interface {
	int64 | float64 | string
}

type Node[T Data] struct {
	Next  *Node[T]
	Value T
}

type Stack[T Data] struct {
	Head *Node[T]
}

func (s *Stack[T]) Push(value T) {
	s.Head = &Node[T]{
		Next:  s.Head,
		Value: value,
	}
}

func (s *Stack[T]) Pop() *Node[T] {
	if s.Head != nil {
		temp := s.Head
		s.Head = s.Head.Next

		return temp
	}

	return nil
}

func main() {
	s := Stack[int64]{}

	s.Push(12)
	s.Push(129)
	s.Push(160)

	fmt.Println(s.Pop().Value)
	fmt.Println(s.Pop().Value)
	fmt.Println(s.Pop().Value)
}

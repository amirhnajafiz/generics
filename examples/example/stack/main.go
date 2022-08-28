// Implementing Stack with Golang generics
// author: Amirhnajafiz
// year: 2022
package main

import "fmt"

// Data Type values.
type Data interface {
	int64 | float64 | string
}

// Node Each stack house is a node.
type Node[T Data] struct {
	Next  *Node[T]
	Value T
}

// Stack Our stack.
type Stack[T Data] struct {
	Head *Node[T]
}

// Push : insert data into stack.
func (s *Stack[T]) Push(value T) {
	s.Head = &Node[T]{
		Next:  s.Head,
		Value: value,
	}
}

// Pop : remove data from stack.
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

package main

import "fmt"

type Data interface {
	int64 | float64 | string
}

type Node[T Data] struct {
	Next  *Node[T]
	Value T
}

func (n *Node[T]) Push(head *Node[T], value T) *Node[T] {
	temp := &Node[T]{
		Next:  head,
		Value: value,
	}

	return temp
}

func main() {
	var head *Node[int64]
	head = head.Push(head, 12)
	head = head.Push(head, 16)

	fmt.Println(head.Value)
}

package main

import "fmt"

type Data interface {
	int64 | float64 | string
}

type Node[T Data] struct {
	Next  *Node[T]
	Value T
}

type LinkedList[T Data] struct {
	Head *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
	temp *Node[T] = l.Head
	node *Node[T] = &Node{
		Next: nil,
		Value: value,
	}

	if temp == nil {
		l.Head = node

		return
	}

	for {
		if temp.Next == nil {
			temp.Next = node

			return
		}
	}
}

// TODO: Remove function
// TODO: Iterate function
func main() {
	l := LinkedList[int64]{}

	l.Add(12)
	l.Add(16)
}

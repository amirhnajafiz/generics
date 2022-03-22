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
	temp := l.Head
	node := &Node[T]{
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

func (l *LinkedList[T]) Remove(value T) {
	temp := l.Head
	prev := l.Head

	for {
		if temp == nil {
			break
		}

		if temp.Value == value {
			if prev == l.Head {
				l.Head = temp.Next
			}
			
			prev.Next = temp.Next
			temp = nil

			return
		}

		prev = temp
		temp = temp.Next
	}
}

func (l *LinkedList[T]) Iterate() {
	temp := l.Head

	for {
		if temp == nil {
			break
		}

		fmt.Print(temp.Value)
		fmt.Print(" ")

		temp = temp.Next
	}
}

func main() {
	l := LinkedList[int64]{}

	l.Add(12)
	l.Add(16)

	l.Remove(12)

	l.Iterate()
}

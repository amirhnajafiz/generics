package main

import "fmt"

type Data interface {
	int64 | float64 | string
}

type Node[T Data] struct {
	Next  *Node[T]
	Value T
}

func main() {
	head := Node[int64]{
		Next:  nil,
		Value: 3,
	}
	fmt.Println(head.Value)
}

package main

import (
	"fmt"
)

type Data interface {
	int | int32 | int64 | float64
}

type Node[T Data] struct {
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
	key    T
}

type Tree[T Data] struct {
	Root *Node[T]
}

func (tree *Tree[T]) Insert(value T) {
	newNode := &Node[T]{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		key:    value,
	}
	if tree.Root == nil {
		tree.Root = newNode
	} else {
		tmp := tree.Root
		for {
			holder := tmp
			if value > tmp.key {
				tmp = tmp.Right
				if tmp == nil {
					holder.Right = newNode
					newNode.Parent = holder
					break
				}
			} else {
				tmp = tmp.Left
				if tmp == nil {
					holder.Left = newNode
					newNode.Parent = holder
					break
				}
			}
		}
	}
}

func (node *Node[T]) Search(key T) *Node[T] {
	if node == nil {
		fmt.Println("KEY DOES NOT EXIST !")
		return nil
	}
	if node.key == key {
		return node
	}
	if key > node.key {
		return node.Right.Search(key)
	} else {
		return node.Left.Search(key)
	}
}

func (root *Node[T]) GetMax() *Node[T] {
	tmp := root
	for {
		holder := tmp
		tmp = tmp.Right
		if tmp == nil {
			return holder
		}
	}
}

func (root *Node[T]) GetMin() *Node[T] {
	tmp := root
	for {
		holder := tmp
		tmp = tmp.Left
		if tmp == nil {
			return holder
		}
	}
}

func (node *Node[T]) GetPredecessor() *Node[T] {
	if node.Left != nil {
		return node.Left.GetMax()
	} else {
		tmp := node
		tmp2 := tmp.Parent
		for tmp2 != nil {
			if tmp != tmp2.Left {
				break
			}
			tmp = tmp2
			tmp2 = tmp2.Parent
		}
		return tmp2
	}
}

func (node *Node[T]) GetSuccessor() *Node[T] {
	if node.Right != nil {
		return node.Right.GetMin()
	} else {
		tmp := node
		tmp2 := tmp.Parent
		for tmp2 != nil {
			if tmp != tmp2.Right {
				break
			}
			tmp = tmp2
			tmp2 = tmp2.Parent
		}
		return tmp2
	}
}

func (tree *Tree[T]) DeleteUseKey(key T) {
	node := tree.Root.Search(key)
	node.DeleteNode()
}

func (node *Node[T]) DeleteNode() bool {
	if node == nil {
		return false
	} else {
		if node.IsLeaf() { //node has no children
			if node.IsRightChild() {
				node.Parent.Right = nil

			} else {
				node.Parent.Left = nil
			}
			node.Parent = nil
		} else if node.Right != nil && node.Left == nil { //node has one children at right
			node.Right.Parent = node.Parent
			if node.IsRightChild() {
				node.Parent.Right = node.Right
			} else {
				node.Parent.Left = node.Right
			}
		} else if node.Left != nil && node.Right == nil { //node has one children at left
			node.Left.Parent = node.Parent
			if node.IsRightChild() {
				node.Parent.Right = node.Left
			} else {
				node.Parent.Left = node.Left
			}
		} else { //node has two children
			sucNod := node.GetSuccessor()
			fmt.Println(sucNod.key)
			holder := sucNod.key
			sucNod.key = node.key
			node.key = holder
			fmt.Println(sucNod.key)
			sucNod.DeleteNode()
		}
		return true
	}
}

func (node *Node[T]) IsLeaf() bool {
	if node.Right == nil && node.Left == nil {
		return true
	}
	return false
}

func (node *Node[T]) IsRightChild() bool {
	if node.key > node.Parent.key {
		return true
	}
	return false
}

func main() {
	tree := Tree[int]{}
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	tree.Insert(40)
}

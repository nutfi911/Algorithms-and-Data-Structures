package bst

import (
	"errors"
)

type Node struct {
	Value     int
	Left      *Node
	Right     *Node
	Frequency int
}

type BST struct {
	Root *Node
}

func New() *BST {
	return &BST{}
}

func (tree *BST) Insert(val int) {
	newNode := &Node{Value: val, Right: nil, Left: nil, Frequency: 0}

	if tree.Root == nil {
		tree.Root = newNode
		return
	}

	current := tree.Root

	for {
		if val == current.Value {
			current.Frequency++
			return
		}

		if val < current.Value {
			if current.Left == nil {
				current.Left = newNode
				break
			}
			current = current.Left
			continue
		}

		if val > current.Value {
			if current.Right == nil {
				current.Right = newNode
				break
			}
			current = current.Right
			continue
		}

	}

}

func (tree *BST) Find(val int) (*Node, error) {
	errNotFound := errors.New("Node not found")

	if tree.Root == nil {
		return &Node{}, errNotFound
	}

	current := tree.Root
	for {
		if val == current.Value {
			return current, nil
		}

		if val < current.Value {
			if current.Left == nil {
				return &Node{}, errNotFound
			}
			current = current.Left
			continue
		}

		if val > current.Value {
			if current.Right == nil {
				return &Node{}, errNotFound
			}
			current = current.Right
			continue
		}
	}
}

func (tree *BST) BFS() []int {
	var queue []*Node
	var result []int

	if tree.Root == nil {
		return result
	}

	queue = append(queue, tree.Root)

	for len(queue) > 0 {
		popped := queue[0]
		queue = append(queue[1:])

		result = append(result, popped.Value)

		if popped.Left != nil {
			queue = append(queue, popped.Left)
		}
		if popped.Right != nil {
			queue = append(queue, popped.Right)
		}
	}

	return result
}

func (node *Node) preOrderTraverse() (result []int) {

	result = append(result, node.Value)

	if node.Left != nil {
		result = append(result, node.Left.preOrderTraverse()...)
	}

	if node.Right != nil {
		result = append(result, node.Right.preOrderTraverse()...)
	}

	return result
}

func (tree *BST) PreOrder() []int {
	var result []int

	if tree.Root == nil {
		return result
	}

	return tree.Root.preOrderTraverse()

}

func (node *Node) postorderTraverse() (result []int) {

	if node.Left != nil {
		result = append(result, node.Left.postorderTraverse()...)
	}

	if node.Right != nil {
		result = append(result, node.Right.postorderTraverse()...)
	}

	result = append(result, node.Value)
	return result
}

func (tree *BST) PostOrder() []int {
	var result []int

	if tree.Root == nil {
		return result
	}

	return tree.Root.postorderTraverse()
}

func (node *Node) inOrderTraverse() (result []int) {

	if node.Left != nil {
		result = append(result, node.Left.inOrderTraverse()...)
	}

	result = append(result, node.Value)

	if node.Right != nil {
		result = append(result, node.Right.inOrderTraverse()...)
	}

	return result
}

func (tree *BST) InOrder() []int {

	var result []int

	if tree.Root == nil {
		return result
	}

	return tree.Root.inOrderTraverse()
}

package bst

import "fmt"

func Run() {

	tree := New()

	// tree.Root = &Node{Value: 10}
	// tree.Root.Right = &Node{Value: 15}
	// tree.Root.Left = &Node{Value: 7}
	// tree.Root.Left.Right = &Node{Value: 9}
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(17)
	// fmt.Println(tree.Root.Value)
	// fmt.Println(tree.Root.Left.Value, tree.Root.Right.Value)
	// fmt.Println(tree.Root.Left.Left.Value)
	fmt.Println(`
			10
		7		15
 	5				17	
	   6
`)

	found, err := tree.Find(5)
	fmt.Println("Find 5: ", found.Value, err)

	fmt.Println("BFS : ", tree.BFS())

	fmt.Println("DFS Pre Order : ", tree.PreOrder())   // 10 7 5 6 15 17
	fmt.Println("DFS Post Order : ", tree.PostOrder()) // 6 5 7 17 15 10
	fmt.Println("DFS Post Order : ", tree.InOrder())   // 5 6 7 10 15 17
}

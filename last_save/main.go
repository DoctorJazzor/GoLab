package main

import "fmt"

type Node struct {
	number int
	left *Node
	right *Node
}

func between (first int, second int, search int) bool{
	if (search <= first && search >= second){
		return true
	} else if (search >= first && search <= second){
		return true
	}
	return false
}

func (Root *Node) Fill_Tree(number int) *Node{
	var Leaf = &Node {number , nil, nil}
	var node **Node
	switch {
		case number < Root.number :
			node = &Root.left
		case number >= Root.number :
			node = &Root.right
	}
	if ((*node) != nil) {
		return (*node).Fill_Tree(number)
	} else {
		(*node) = Leaf
	}
	return Root
}

func (Root *Node) Fill_Tree_A(number int) *Node{
	var Leaf = &Node {number , nil, nil}
	switch {
		case number < Root.number :
			if (Root.left != nil){
				if (Root.left.number < number){
					Leaf.left = Root.left
					Leaf.right = Root
					Root = Leaf
					return Root
				} else {
					return Root.left.Fill_Tree(number)
				}
			} else {
				Root.left = Leaf
			}
		case number >= Root.number :
			if (Root.right != nil){
				if (Root.right.number >= number){
					Leaf.left = Root
					Leaf.right = Root.right
					Root = Leaf
					return Root
				} else {
					return Root.right.Fill_Tree(number)
				}
			} else {
				Root.right = Leaf
			}
	}
	return Root
}


var count int
func (Root Node) read() Node{
	fmt.Printf("Number is %d for deep in %d\n",  Root.number, count)
	if (Root.left != nil){
		count++
		fmt.Printf(" Left ")
		Root.left.read()
	} else {
		fmt.Printf(" STOP LEAF left\n")
	}
	if (Root.right != nil) {
		count++
		fmt.Printf(" Right ")
		Root.right.read()
	} else {
		fmt.Printf(" STOP LEAF right\n")
	}
	count --
	return Root
}

func main () {
	var root Node
	root.number = 7
	/*for inc := 1; inc < 17; inc ++ {
		root.Fill_Tree(inc)
	}*/
	root.Fill_Tree(3)
	root.Fill_Tree(10)
	root.Fill_Tree(8)
	root.Fill_Tree(9)
	root.Fill_Tree(1)
	root.Fill_Tree(4)
	root.Fill_Tree(11)
	root.read()
	fmt.Printf("Hello World!\n")
}

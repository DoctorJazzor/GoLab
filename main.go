package main

import (
		"fmt"
		"strings"
	)

type Node struct {
	number int
	master *Node
	left *Node
	right *Node
}

func (Root *Node) Fill_Tree_A(number int) *Node{
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
		(*node) = new (Node)
		(*node).number = number
	}
	return Root
}

func (Root *Node) fill_infixe (number string) *Node{
	if (number is not number){
		if (number == "+"){
			Root.number = number
		}
		switch {
			case Root.left == nil:
				Root.left.fill_infixe(number)
			case Root.right == nil:
				Root.left.fill_infixe(number)
		}
	} else {
		Root.number = number
	}
}

func (Root *Node) fill_infixe (number string) *Node{
	if (number is not number){
		if (Root.left == nil){
			Root.left.fill_infixe(number)
		}
		Root.number = number
		if (Root.right == nil){
			Root.left.fill_infixe(number)
		}
	} else {
		Root.number = number
	}
}

func (Root *Node) fill_infixe (number string) *Node{
	if (number is number){
		if (Root.left == nil){
			Root.left = new (*Node)
			Root
			Root.left.fill_infixe(number)
		}
		Root.number = number
		if (Root.right == nil){
			Root.left.fill_infixe(number)
		}
	} else {
		Root.number = number
	}
}

func calcul (operation string) int{
	var root Node
	var char []string
	char = strings.Split(operation, " ")
	for inc := 0; inc < len(char); inc++{
		if (char [inc] != ")"){
			return 1
		} else if (char [inc] == "+"){
			return 0
		} else {
			return -1
		}
	}
	return root
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
	calcul("a a+ a")
	var root Node
	root.number = 7
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

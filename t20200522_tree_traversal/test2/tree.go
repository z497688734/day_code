package main

import "fmt"

/**
找到最左下的节点
 */
type BinaryTree struct {
	Data  int
	Left  *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data int) *BinaryTree {
	return &BinaryTree{Data: data}
}



func getLeftButtomNodeVal(root *BinaryTree,maxDepth *int,val *int,depth int){
	if root == nil{
		return
	}

	getLeftButtomNodeVal(root.Left,maxDepth,val,depth+1)
	getLeftButtomNodeVal(root.Right,maxDepth,val,depth+1)
	if depth > *maxDepth{
		*maxDepth = depth
		*val = root.Data
		fmt.Println("maxDepth",*maxDepth,"root.Data",root.Data)

	}

}


func getLeftButtomNodeVal2(root *BinaryTree,maxDepth int,val *int,depth int){
	if root == nil{
		return
	}

	getLeftButtomNodeVal2(root.Left,maxDepth,val,depth+1)
	getLeftButtomNodeVal2(root.Right,maxDepth,val,depth+1)
	if depth > maxDepth{
		fmt.Println("maxDepth",maxDepth,"root.Data",root.Data)
		maxDepth = depth
		*val = root.Data

	}

}


func main() {
	t := NewBinaryTree(1)
	t.Left  = NewBinaryTree(2)
	t.Right = NewBinaryTree(3)
	t.Left.Left = NewBinaryTree(4)
	t.Left.Right = NewBinaryTree(5)
	t.Right.Left = NewBinaryTree(6)
	t.Right.Right = NewBinaryTree(7)
	//t.Right.Right.Left = NewBinaryTree(8)
	t.Left.Left.Left = NewBinaryTree(8)



	var maxDepth = 0
	var val  = 0
	//getLeftButtomNodeVal(t,&maxDepth,&val,0)
	getLeftButtomNodeVal(t,&maxDepth,&val,0)

	fmt.Println(maxDepth)
	fmt.Println(val)
}

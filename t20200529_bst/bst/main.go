package main

import "fmt"

type Node struct {
	Key int
	Value string
	Num int //以该节点为跟节点树的节点数
	Left *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (tree  *Tree) Get(key int) (found bool, value string)  {
	return   get(tree.root,key)
}

func getSize( node *Node)  int {
	if node == nil {
		return  0
	}
	return node.Num
}
func get(node *Node,key int) (found bool, value string)  {
	if node == nil{
		return  false,"";
	}
	if  key < node.Key {
		return  get(node.Left,key)
	}else if key > node.Key{
		return get(node.Right,key)
	}else{
		return true ,node.Value
	}
}

func (tree *Tree)Put(key int,value string)  {
	tree.root = put(tree.root,key,value)

}

func put(node *Node,key int , value string) *Node  {
	if node == nil{
		return  &Node{Key:key,Value:value,Num:1}
	}
	if  key < node.Key {
		node.Left = put(node.Left,key,value)
	}else if key > node.Key{
		node.Right = put(node.Right,key,value)
	}else{
		node.Value = value
	}


	node.Num = getSize(node.Left) + getSize(node.Right) + 1
	return  node
}

func (tree *Tree) printTree(node *Node)  {
	fmt.Printf("%+v",node)
	if node.Left != nil{
		tree.printTree(node.Left)
	}

	if node.Right != nil{
		tree.printTree(node.Right)
	}

}


func main()  {
	var tree Tree =  Tree{nil}
	tree.Put(4,"4")
	tree.Put(5,"5")
	tree.Put(3,"3")
	tree.Put(6,"6")
	tree.Put(1,"1")


	tree.printTree(tree.root)

}
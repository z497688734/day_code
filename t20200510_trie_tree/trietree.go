package main

import "fmt"

type TrieNode struct {
	children map[interface{}] *TrieNode
	isEnd bool
}

func newTrieNode() *TrieNode {
	return  &TrieNode{children:make(map[interface{}]*TrieNode),isEnd:false}
}


type TrieTree struct {
	root *TrieNode
}

func newTrieTree() *TrieTree  {
	return &TrieTree{root:newTrieNode()}
}


func (tree *TrieTree) insert(word []interface{})  {
	node := tree.root
	for i:=0;i< len(word);i++ {
		_,ok := node.children[word[i]]
		if !ok{
			node.children[word[i]] =  newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

func (tree *TrieTree)search(word[]interface{}) bool {
	node := tree.root
	for i:=0;i<len(word);i++ {
		_,ok := node.children[word[i]]
		if !ok{
			return false
		}
		node = node.children[word[i]]
	}
	return  node.isEnd
}

func (tree *TrieTree) printTree(node *TrieNode)  {
	fmt.Printf("%+v",node.children)
	if len(node.children) == 0{
		return
	}
	for k,v := range node.children{
		fmt.Println(v)
		tree.printTree(node.children[k])
	}
}


func  main()  {
	a := "abc"
	tree := newTrieTree()
	tree.insert(tranKey(a))
	exist := tree.search(tranKey(a))
	fmt.Println(exist)


	tree.insert(tranKey("abcd"))

	fmt.Println("begin print tree")
	tree.printTree(tree.root)
	fmt.Println("end print tree")

}

func tranKey(key string) []interface{}  {
	var newKey []interface{} = make([]interface{},len(key))

	for i:=0;i<len(key);i++ {
		newKey[i] = (interface{})(key[i])
	}
	return newKey
}
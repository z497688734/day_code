package main

import (
	"container/list"
	"fmt"
)

// Binary Tree
type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

// 先序遍历-非递归
func (bt *BinaryTree) PreOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			res = append(res, t.Data)//visit
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 中序遍历-非递归
func (bt *BinaryTree) InOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			res = append(res, t.Data)//visit
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 后序遍历-非递归
func (bt *BinaryTree) PostOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	var preVisited *BinaryTree

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}

		v   := stack.Back()
		top := v.Value.(*BinaryTree)

		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right{
			res = append(res, top.Data)//visit
			preVisited = top
			stack.Remove(v)
		}else {
			t = top.Right
		}
	}
	return res
}



func findLeftButtomNodeVal(root *BinaryTree,maxDepth *int,val *interface{},depth int){
	if root == nil{
		return
	}

	findLeftButtomNodeVal(root.Left,maxDepth,val,depth+1)
	findLeftButtomNodeVal(root.Right,maxDepth,val,depth+1)
	if depth > *maxDepth{
		*maxDepth = depth
		*val = root.Data
	}
}


func findLeftButtomNodeValByBFS(root *BinaryTree)(val interface{}){
	res := make([]interface{}, 0)
	if root == nil {
		return 0
	}
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		length := len(queue)
		var isSet = false
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				if isSet == false{
					val = queue[0].Left.Data
					isSet = true
				}
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
				if isSet == false{
					val = queue[0].Right.Data
					isSet = true
				}
			}
			res = append(res, queue[0].Data)
			queue = queue[1:]
		}
	}

	return
}

//宽度遍历
func BFS(root *BinaryTree) []interface{}  {
	res := make([]interface{}, 0)
	if root == nil {
		return res
	}
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		res = append(res, queue[0].Data)
		queue = queue[1:]
	}
	return  res
}

//最大深度
func getMaxDepthByBFS(root *BinaryTree)  (depth int) {
	res := make([]interface{}, 0)
	if root == nil {
		return 0
	}
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		length := len(queue)
		depth ++
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res = append(res, queue[0].Data)
			queue = queue[1:]
		}
	}
	return  depth
}


//最大深度
func getMaxDepthByRecursion(root *BinaryTree)  (depth int) {
	if root == nil{
		return  0
	}
	leftMaxDepth := getMaxDepthByRecursion(root.Left)
	rightMaxDepth := getMaxDepthByRecursion(root.Right)

	return  1 + max(leftMaxDepth, rightMaxDepth)
}

//最小深度
func getMinDepthByRecursion(root *BinaryTree)  (depth int) {
	if root == nil{
		return  0
	}
	leftMinDepth := getMinDepthByRecursion(root.Left)
	rightMinDepth := getMinDepthByRecursion(root.Right)

	return  1 + min(leftMinDepth, rightMinDepth)
}


func findRightButtomNodeVal(root *BinaryTree,maxDepth *int,val *interface{},depth int){
	if root == nil{
		return
	}
	findRightButtomNodeVal(root.Right,maxDepth,val,depth+1)
	findRightButtomNodeVal(root.Left,maxDepth,val,depth+1)
	if depth > *maxDepth{
		*maxDepth = depth
		*val = root.Data
	}
}

func findRightButtomNodeVal2(root *BinaryTree,val *interface{},depth int) (maxDepth int){
	if root == nil{
		return 0
	}
	leftMaxDepth := findRightButtomNodeVal2(root.Left,val,depth+1)
	rightMaxDepth := findRightButtomNodeVal2(root.Right,val,depth+1)
	maxDepth = 1 + max(leftMaxDepth, rightMaxDepth)

	fmt.Println(depth,maxDepth,root.Data)
	if depth > maxDepth{
		maxDepth = depth
		*val = root.Data
	}
	return  maxDepth
}





//反转树的左右节点
func swapLeftAndRight(root *BinaryTree)  {
	if root == nil{
		return
	}
	swapLeftAndRight(root.Left)
	swapLeftAndRight(root.Right)

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
}
func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func main() {
	t := NewBinaryTree(1)
	t.Left  = NewBinaryTree(2)
	t.Right = NewBinaryTree(3)
	//t.Left.Left = NewBinaryTree(4)
	//t.Left.Right = NewBinaryTree(5)
	t.Right.Left = NewBinaryTree(6)
	t.Right.Right = NewBinaryTree(7)
	t.Right.Right.Left = NewBinaryTree(8)
	t.Right.Right.Right = NewBinaryTree(9)


	//fmt.Println(t.PreOrderNoRecursion())
	//fmt.Println(t.InOrderNoRecursion())
	//fmt.Println(t.PostOrderNoRecursion())

	//var maxDepth = 1
	//var val  interface{}
	//findLeftButtomNodeVal(t,&maxDepth,&val,1)
	//fmt.Println(val)

	//res  := getMaxDepthByRecursion(t)
	//res := getMinDepthByRecursion(t)

	//res := findLeftButtomNodeValByBFS(t)
	//fmt.Println(res)

	//swapLeftAndRight(t)
	//res := BFS(t)
	//fmt.Println(res)

	//var maxDepth = 1
	//var val  interface{}
	//findRightButtomNodeVal(t,&maxDepth,&val,1)
	//fmt.Println(val)


	var val  interface{}
	findRightButtomNodeVal2(t,&val,1)
	fmt.Println(val)


}

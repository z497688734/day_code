package main

import (
	"fmt"
)

type Node struct {
	Next	*Node
	Data     int
}





func (head  *Node) AddLast(v int){
	var last *Node = head
	for last.Next != nil {
		last = last.Next
	}

	node := &Node{Next:nil,Data:v}
	last.Next = node
}


func (head *Node) Revert()(*Node){
	cur := head
	var preNode *Node
	for cur != nil {
		nextNode := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}
	return preNode
}




func RevertRecursion(cur *Node)  (*Node) {
	if cur.Next == nil {
		return cur
	}

	newhead := RevertRecursion(cur.Next)
	cur.Next.Next = cur
	cur.Next = nil
	return newhead

}
func (head  *Node) Print( )  {
	cur := head
	for {
		fmt.Println("print####",cur.Data)
		cur = cur.Next
		if(cur == nil){
			break
		}
	}
}

func main()  {
	head := &Node{Data:1,Next:nil}
	head.AddLast(2)
	head.AddLast(3)
	head.AddLast(4)
	head.AddLast(5)

	head.Print()
	fmt.Println("begin revert....")
	rHead:= head.Revert()
	rHead.Print()
	fmt.Println("begin revert....")

	r2 := RevertRecursion(rHead)
	r2.Print()
	//fmt.Println("end revert....")
	//
	//rLink.Print()

}
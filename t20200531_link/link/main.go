package main

import (
	"fmt"
	"runtime/debug"
)

type Node struct {
	Next	*Node
	Data     int
}



type Link struct {
	Head *Node
}


func (link  *Link) Put(v int){
	node := &Node{Next:nil,Data:v}
	if link.Head == nil{
		link.Head = node
	}else{
		node.Next = link.Head
		link.Head = node
	}

}


//func (link *Link) Revert()(*Link){
//	cur := link.Head
//	var preNode *Node
//	for cur != nil {
//		nextNode := cur.Next
//		cur.Next = preNode
//		preNode = cur
//		cur = nextNode
//	}
//	return &Link{Head:preNode}
//}


func (link *Link) Revert()(*Link){
	head := revert(link.Head)
	return  &Link{Head :head}
}

func revert(cur *Node)  (*Node) {
	fmt.Println("begin print debug...")
	fmt.Printf("%s", debug.Stack())
	fmt.Println("end print end...")

	if cur.Next == nil {
		return cur
	}

	newhead := revert(cur.Next)
	cur.Next.Next = cur
	cur.Next = nil
	return newhead

}
func (link  *Link) Print( )  {
	cur := link.Head
	for {
		fmt.Println("print####",cur.Data)
		cur = cur.Next
		if(cur == nil){
			break
		}
	}
}

func main()  {
	link := &Link{}
	link.Put(1)
	link.Put(2)
	link.Put(3)
	link.Put(4)
	link.Put(5)

	link.Print()



	fmt.Println("begin revert....")
	rLink:= link.Revert()
	fmt.Println("end revert....")

	rLink.Print()

}
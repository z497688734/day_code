package main

import "fmt"

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


func (link *Link) Revert()(*Link){
	cur := link.Head
	var preNode *Node = nil
	for cur != nil {
		nextNode := cur.Next
		cur.Next  = preNode
		preNode = cur
		cur = nextNode
	}
	return &Link{Head:preNode}
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
	link.Print()



	fmt.Println("begin revert....")
	rLink:= link.Revert()
	fmt.Println("end revert....")

	rLink.Print()

}
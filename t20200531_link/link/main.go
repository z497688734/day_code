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
	fmt.Println("newHead",newhead)
	cur.Next.Next = cur
	cur.Next = nil
	return newhead

}

func RevertRecursionKGroup(head *Node,num int)  (*Node) {
	var  nextNode,preNode *Node
	//下一组开始
	curGroupStart,nextGroupStart := head,head

	for i:=0;i<num;i++ {
		if nextGroupStart != nil{
			nextGroupStart = nextGroupStart.Next
		}else{
			return  head
		}
	}

	index :=0
	for curGroupStart != nil && index < num{
		nextNode = curGroupStart.Next
		curGroupStart.Next = preNode
		preNode = curGroupStart
		curGroupStart = nextNode
		index++
	}

	fmt.Println(".....")
	fmt.Println(nextNode,nextGroupStart)
	fmt.Println(".....")

	head.Next = RevertRecursionKGroup(nextNode,2)
	fmt.Println("header",head)
	return  preNode

}

func reverseKGroup(head *Node, k int) *Node {
	var pre, next *Node
	var cur, p *Node = head, head
	groupSize, index := 0, 0
	// 将链表的前k个结点组成一组
	for p != nil && groupSize < k {
		p = p.Next
		groupSize++
	}
	if groupSize == k {
		// 翻转该k个结点组成的子链表
		for cur != nil && index < k {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			index++
		}
		// 继续递归地执行上述过程
		if next != nil {
			// head指向子链表翻转后的尾节点
			head.Next = reverseKGroup(next, k)
		}
		// pre记录了为子链表翻转后的头结点
		return pre
	} else {
		// 子链表长度不足k，不翻转
		return head
	}
}


func (head  *Node) Print( )  {
	cur := head
	for {
		cur = cur.Next
		if cur == nil{
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
	head.AddLast(6)
	head.AddLast(7)
	head.AddLast(8)
	head.AddLast(9)
	head.AddLast(10)

	//head.Print()
	//fmt.Println("begin revert....")
	//rHead:= head.Revert()
	//rHead.Print()
	//fmt.Println("begin revert....")

	//r2 := RevertRecursion(head)
	//r2.Print()
	//fmt.Println("end revert....")
	//
	//rLink.Print()

	a := RevertRecursionKGroup(head,2)
	a.Print()

	 //b := reverseKGroup(head,2)
	 //b.Print()
}
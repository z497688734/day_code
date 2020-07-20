package main

import (
	"fmt"
	"runtime"
	"time"
)

type Node struct {
	Next	*Node
	Data     int
}

func  New() ( *Node)  {
	head := &Node{}
	fmt.Println(head.Next)
	return head
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

func (head *Node)ReverRecursion2() (*Node)  {
	if head.Next == nil{
		return  head
	}

	newhead := head.Next.ReverRecursion2()
	head.Next.Next = head
	head.Next = nil
	return newhead
}

func ReverRecursion3(pre *Node ,cur *Node) (*Node)  {
	if cur == nil{
		return  nil
	}
	if cur.Next == nil {
		cur.Next = pre
		return cur
	}

	next := cur.Next
	cur.Next = pre
	return ReverRecursion3(cur,next);
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
		if cur == nil{
			break
		}
		cur = cur.Next
	}
}

func main()  {
	go func() {
		for {
			var m runtime.MemStats
			tick := time.After(1 * time.Second)
			select {
			case <-tick:
				runtime.ReadMemStats(&m)
				gb := 1024 * 1024 * 1024.0
				logstr := fmt.Sprintf("\nAlloc = %v\tTotalAlloc = %v\tSys = %v\t NumGC = %v\n", float64(m.Alloc)/gb, float64(m.TotalAlloc)/gb, float64(m.Sys)/gb, m.NumGC)
				fmt.Println(logstr)
			}
		}
	}()

	head := New()
	for j:=1;j<100000;j++ {
		head.AddLast(j)
	}
	fmt.Println("begin ReverRecursion3 \n")
	for  i:=0 ;i<10;i++{
		time.Sleep(1 * time.Second)
		ReverRecursion3(nil,head)
	}
	fmt.Println("end ReverRecursion3\n")

	fmt.Println("begin ReverRecursion2\n")
	for  i:=0 ;i<10;i++{
		time.Sleep(1 * time.Second)
		head.ReverRecursion2()
	}
	fmt.Println("end ReverRecursion2\n")
	select {
	}


	//head.Print()
	//fmt.Println("begin revert....")
	//rHead:= head.ReverRecursion2()
	//rHead.Print()
	//fmt.Println("begin revert....")

	//r2 := RevertRecursion(head)
	//r2.Print()
	//fmt.Println("end revert....")
	//
	//rLink.Print()


	//a := RevertRecursionKGroup(head,2)
	//a := ReverRecursion3(nil,head)
	//a.Print()

	 //b := reverseKGroup(head,2)
	 //b.Print()
}
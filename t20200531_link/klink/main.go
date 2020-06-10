package main
import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var pre, next *ListNode
	var cur, p *ListNode = head, head
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

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l := &ListNode {a[0], nil}
	p := l
	for _, r := range a[1:] {
		node := &ListNode {r, nil}
		p.Next = node
		p = p.Next
	}

	fmt.Println(l)
	fmt.Println(p)
	fmt.Println("before reverse...")
	cur := l
	for cur != nil {
		if cur.Next != nil {
			fmt.Printf("%d->", cur.Val)
		} else {
			fmt.Printf("%d", cur.Val)
		}
		cur = cur.Next
	}
	fmt.Println("end reverse...")

	k := 3
	res := reverseKGroup(l, k)
	for res != nil {
		if res.Next != nil {
			fmt.Printf("%d->", res.Val)
		} else {
			fmt.Printf("%d", res.Val)
		}
		res = res.Next
	}
}


package main

import "fmt"

/*link node reverse*/

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func reverse(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	var pre *LinkNode
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func main() {
	head := &LinkNode{}
	//cur := &LinkNode{}
	cur := head
	for i := 1; i <= 5; i++ {
		cur.Next = &LinkNode{Val: i}
		cur = cur.Next
	}

	head1 := head.Next
	cur1 := head1
	for cur1 != nil {
		fmt.Println(cur1.Val)
		cur1 = cur1.Next
	}

	res := reverse(head1)

	//print after
	cur2 := res
	for cur2 != nil {
		fmt.Println(cur2.Val)
		cur2 = cur2.Next
	}

}

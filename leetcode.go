package main

import "fmt"

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3
	print(l1)

	fmt.Print("\n")

	l2 := new(ListNode)
	l2.Val = 3
	l2.Next = new(ListNode)
	l2.Next.Val = 2
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 1
	print(l2)

	fmt.Print("\n")

	print(addTwoNumbers(l1, l2))

}

func print(node *ListNode)  {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next;
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode);

	temp := result;

	for l1 != nil && l2 != nil  {
		Val := l1.Val + l2.Val
		addResult(temp, Val)
		temp = temp.Next;
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil   {
		Val := l1.Val
		addResult(temp, Val)
		temp = temp.Next;
		l1 = l1.Next
	}

	for l2 != nil   {
		Val := l2.Val
		addResult(temp, Val)
		temp = temp.Next;
		l2 = l2.Next
	}

	return result.Next;
}

func addResult(Node *ListNode, Val int)  {
	if Node.Next == nil {
		Node.Next = new(ListNode)
		Node =  Node.Next
		if Val < 10 {
			lessThanTen(Node, Val)
		}else {
			moreThanTen(Node, Val)
		}
	}else {
		Node = Node.Next
		Val = Val + Node.Val
		if Val < 10 {
			lessThanTen(Node, Val)
		}else {
			moreThanTen(Node, Val)
		}
	}
}

func lessThanTen(Node *ListNode, Val int)  {
	Node.Val = Val;
}

func moreThanTen(Node *ListNode, Val int)  {
	Node.Next = new(ListNode)
	Node.Next.Val = 1;
	Node.Val = Val % 10;
}

type ListNode struct {
	Next *ListNode
	Val int
}

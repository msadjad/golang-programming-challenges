package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	first, second := makeLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), makeLinkedList([]int{9, 9, 9, 9})
	printLinkedList(first)
	printLinkedList(second)
	result := addTwoNumbers(first, second)
	printLinkedList(result)
}

func makeLinkedList(list []int) *ListNode {
	var linkedList *ListNode = nil

	if len(list) == 0 {
		return linkedList
	}

	linkedList = new(ListNode)
	current := linkedList

	for i := 0; i < len(list); i++ {
		current.Val = list[i]
		if i != len(list)-1 {
			current.Next = new(ListNode)
			current = current.Next
		}
	}

	return linkedList
}

func printLinkedList(list *ListNode) {
	for current := list; current != nil; current = current.Next {
		fmt.Printf("%v ", current.Val)
	}
	fmt.Print("\n")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode = nil
	var carry int = 0
	l1Current, l2Current := l1, l2

	if l1Current == nil && l2Current == nil {
		return result
	}

	result = new(ListNode)
	lCurrent := result

	for ; l1Current != nil && l2Current != nil; l1Current, l2Current = l1Current.Next, l2Current.Next {
		lCurrent.Val = (l1Current.Val + l2Current.Val + carry) % 10
		carry = (l1Current.Val + l2Current.Val + carry) / 10

		if l1Current.Next != nil || l2Current.Next != nil {
			lCurrent.Next = new(ListNode)
			lCurrent = lCurrent.Next
		}
	}

	for ; l1Current != nil; l1Current = l1Current.Next {
		lCurrent.Val = (l1Current.Val + carry) % 10
		carry = (l1Current.Val + carry) / 10

		if l1Current.Next != nil {
			lCurrent.Next = new(ListNode)
			lCurrent = lCurrent.Next
		}
	}

	for ; l2Current != nil; l2Current = l2Current.Next {
		lCurrent.Val = (l2Current.Val + carry) % 10
		carry = (l2Current.Val + carry) / 10

		if l2Current.Next != nil {
			lCurrent.Next = new(ListNode)
			lCurrent = lCurrent.Next
		}
	}

	if carry == 1 {
		lCurrent.Next = new(ListNode)
		lCurrent = lCurrent.Next
		lCurrent.Val = 1
	}

	return result
}

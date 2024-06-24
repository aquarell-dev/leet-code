package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}

// helper function to print the linked list
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var vc1, vc2, sum int
	res := &ListNode{Val: 0, Next: nil}
	cur := res
	carry := 0
	current1 := l1
	current2 := l2

	for {
		if current1 == nil && current2 == nil {
			if carry == 1 {
				cur.Next = &ListNode{Val: 1, Next: nil}
			}
			return res.Next
		}

		if current1 != nil {
			vc1 = current1.Val
			current1 = current1.Next
		} else {
			vc1 = 0
		}

		if current2 != nil {
			vc2 = current2.Val
			current2 = current2.Next
		} else {
			vc2 = 0
		}

		sum = vc1 + vc2 + carry

		if sum >= 10 {
			carry = 1
			sum = sum - 10
		} else {
			carry = 0
		}

		cur.Next = &ListNode{Val: sum, Next: nil}
		cur = cur.Next
	}
}

func main() {
	// Create linked lists l1 and l2
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4, 4})

	res := addTwoNumbers(l1, l2)

	// Print the linked lists
	fmt.Print("l1: ")
	printLinkedList(l1)
	fmt.Print("l2: ")
	printLinkedList(l2)
	fmt.Print("res: ")
	printLinkedList(res)
}

package main

import "fmt"

func main() {
	b := addTwoNumbers(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	})
	fmt.Println(b)

}

func twoSum(nums []int, target int) []int {
	ra := 0
	rb := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				ra = i
				rb = j
			}
		}
	}
	result := []int{ra, rb}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	var add int
	for l1 != nil || l2 != nil {
		q := 0
		w := 0
		if l1.Val != 0 {
			q = l1.Val
		}
		if l2.Val != 0 {
			w = l2.Val
		}
		var sum int
		sum = add + q + w
		add = sum / 10

		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if add > 0 {
		cur.Next = &ListNode{
			Val: add,
		}
	}

	return dummyHead.Next

}

func testListNode() {
	Head := &ListNode{}
	curr := Head

	q := 1
	curr.Next = &ListNode{Val: q}
	curr = curr.Next
	fmt.Println("--New--")
	fmt.Println(Head)
	fmt.Println(Head.Next)

	w := 2
	curr.Next = &ListNode{Val: w}
	curr = curr.Next
	fmt.Println("--New--")
	fmt.Println(Head)
	fmt.Println(Head.Next)
	fmt.Println(Head.Next.Next)

	e := 3
	curr.Next = &ListNode{Val: e}
	curr = curr.Next
	fmt.Println("--New--")
	fmt.Println(Head)
	fmt.Println(Head.Next)
	fmt.Println(Head.Next.Next)
}

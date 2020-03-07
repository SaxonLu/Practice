package main

import (
	"fmt"
	"sync"
)

func main() {

	// q := twoSum([]int{57, 92, 2, 5, 7, 22, 85}, 9)
	// fmt.Print(q)
	// b := addTwoNumbers(&ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 4,
	// 		Next: &ListNode{
	// 			Val:  3,
	// 			Next: nil,
	// 		},
	// 	},
	// }, &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val:  9,
	// 			Next: nil,
	// 		},
	// 	},
	// })
	b := reverse(120)
	fmt.Println(b)
	c := isPalindrome(1441)
	fmt.Print(c)
}

// 傳入int陣列 找出陣列內相符兩個int 其和等於目標和
// 傳入 1 4 6 找出 想加為7的兩個陣列元素
// ->兩個陣列互相比較
//
//	4 6
//  1 4 6

func twoSum(nums []int, target int) []int {
	ra := 0
	rb := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[j]+nums[i] {
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

// 連結串列(鍊式資料)
// 給定兩個非空鍊表來表示兩個非負整數。位數按照逆序方式存儲，它們的每個節點只存儲單個數字。將兩數相加返回一個新的鍊表。
// 你可以假設除了數字0 之外，這兩個數字都不會以零開頭。
//
// 迴圈讀鍊表資料
// 加總 > 10的溢位處理
// 寫入新的鍊表(運用到指標型態的堆疊)
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

// 觀察指標型態堆疊用
// --1--
// &{0 0x40c140}
// &{1 <nil>}
// <nil>
// --2--
// &{0 0x40c140}
// &{1 0x40c160}
// &{2 <nil>}
// <nil>
// --3--
// &{0 0x40c140}
// &{1 0x40c160}
// &{2 0x40c180}
// &{3 <nil>}
func testListNode() {
	Head := &ListNode{}
	curr := Head

	curr.Next = &ListNode{Val: 1}
	curr = curr.Next
	fmt.Println("--1--")
	fmt.Println(Head)
	fmt.Println(Head.Next)

	curr.Next = &ListNode{Val: 2}
	curr = curr.Next
	fmt.Println("--2--")
	fmt.Println(Head)
	fmt.Println(Head.Next)
	fmt.Println(Head.Next.Next)

	curr.Next = &ListNode{Val: 3}
	curr = curr.Next
	fmt.Println("--3--")
	fmt.Println(Head)
	fmt.Println(Head.Next)
	fmt.Println(Head.Next.Next)
}

// 翻轉數字
// 123 to 321
// -123 to -321
// 120 to 21
// 題目有規範 範圍為 −2^31,  2^31 − 1
// 如果溢出則回傳 0
//
// 每次取得尾數 %10
// 然後將尾數 *10堆疊成結果
// 因每次取走尾數(個位)
// 所以原數字需除10
// 範圍作檢核
func reverse(x int) int {
	var nums, newnums int
	// 直到X等於0 break
	for x != 0 {
		a := x % 10
		newnums = nums*10 + a

		nums = newnums
		x /= 10

		// 範圍是 −2^31,  2^31 − 1]
		MaxInt32 := 1<<31 - 1
		MinInt32 := -1 << 31

		if nums > MaxInt32 || nums < MinInt32 {
			return 0
		}
	}

	return nums
}

// 左/右移運算符
// int a = 5;        // binary: 00000000000000000000000000000101
// int b = a << 3;   // binary: 00000000000000000000000000101000, 等於十進制的 40
// int c = b >> 3;   // binary: 00000000000000000000000000000101, 回到一開始的 5

// 判斷傳入數字是否為回文格式
// 12321 yes
// -12321 NO -> 12321-
// 10 NO -> 01
// 判別是否為負數，其尾數是否為 0
// 迴圈如果處理至翻轉數 > 原數時 = 已處理過一半 所以可判別是否為回文
// 取個位數 並 * 10方式 疊加給新數
// 最後判別原數(以取一半)是否和新數相等(回文)
// 額外判斷 奇數狀態 EX 12321 去除中間一位來判別
// /= 除法指派

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	temp := 0
	for x > temp {
		temp = temp*10 + x%10
		x /= 10
	}
	return x == temp || x == temp/10

}

/// 單例模式 練習

var m *Manager
var once *sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

func (p Manager) Manage() {
	fmt.Println("Call Manage")
}

type Manager struct{}

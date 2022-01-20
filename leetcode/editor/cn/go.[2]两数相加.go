/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-20 22:28:19
 * @version     v1.0
 * @filename    go.[2]两数相加
 * @description
 ***************************************************************************/
package main

import (
	"math"
)

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学 👍 7385 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	dummy := new(ListNode)
	dummy.Next = ret
	sum := 0
	carry := 0
	for {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		ret.Val = sum % 10
		carry = int(math.Floor(float64(sum / 10)))
		if l1 == nil && l2 == nil {
			break
		}
		ret.Next = new(ListNode)
		ret = ret.Next
	}
	if carry > 0 {
		ret.Next = new(ListNode)
		ret.Next.Val = 1
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	l1 := new(ListNode)
//	l1.Val = 9
//	l1.Next = new(ListNode)
//	l1.Next.Val = 9
//	l1.Next.Next = new(ListNode)
//	l1.Next.Next.Val = 9
//
//	l2 := new(ListNode)
//	l2.Val = 9
//	l2.Next = new(ListNode)
//	l2.Next.Val = 9
//	l2.Next.Next = new(ListNode)
//	l2.Next.Next.Val = 9
//	l2.Next.Next.Next = new(ListNode)
//	l2.Next.Next.Next.Val = 9
//
//	ret := addTwoNumbers(l1, l2)
//	fmt.Println(ret.Next.Next.Next.Val)
//}

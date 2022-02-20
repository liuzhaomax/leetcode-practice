/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-21 00:20:35
 * @version     v1.0
 * @filename    go.[148]排序链表
 * @description
 ***************************************************************************/
package main

import (
	"sort"
)

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
// Related Topics 链表 双指针 分治 排序 归并排序 👍 1472 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dump := new(ListNode)
	dump.Next = head

	values := make([]int, 0)
	cur := dump.Next
	for {
		if cur != nil {
			values = append(values, cur.Val)
		} else {
			break
		}
		cur = cur.Next
	}
	sort.Ints(values)
	cur = dump
	for _, v := range values {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return dump.Next
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	n4 := &ListNode{Val: 3, Next: nil}
//	n3 := &ListNode{Val: 1, Next: n4}
//	n2 := &ListNode{Val: 2, Next: n3}
//	n1 := &ListNode{Val: 4, Next: n2}
//	res := sortList(n1)
//	cur := res
//	for {
//		fmt.Println(cur.Val)
//		if cur.Next == nil {
//			break
//		}
//		cur = cur.Next
//	}
//}

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-25 15:08:59
 * @version     v1.0
 * @filename    go.[19]删除链表的倒数第 N 个结点
 * @description
 ***************************************************************************/
package main

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针 👍 1769 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	curr := dummy
	prev := dummy
	for i := 0; i < n; i++ {
		curr = curr.Next
	}
	for {
		if curr.Next == nil {
			prev.Next = prev.Next.Next
			break
		}
		prev = prev.Next
		curr = curr.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

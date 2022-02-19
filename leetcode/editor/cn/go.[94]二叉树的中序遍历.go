/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 02:24:35
 * @version     v1.0
 * @filename    go.[94]二叉树的中序遍历
 * @description
 ***************************************************************************/
package main

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[2,1]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1272 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res = []int{}
	traverse(root, &res)
	return res
}

func traverse(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, res)
	*res = append(*res, node.Val)
	traverse(node.Right, res)
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	n3 := &TreeNode{3, nil, nil}
//	n2 := &TreeNode{2, n3, nil}
//	n1 := &TreeNode{1, nil, n2}
//	fmt.Println(inorderTraversal(n1))
//}

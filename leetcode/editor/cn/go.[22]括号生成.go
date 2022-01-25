/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-25 16:07:16
 * @version     v1.0
 * @filename    go.[22]括号生成
 * @description
 ***************************************************************************/
package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 动态规划 回溯 👍 2321 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	var arr []string
	type Node struct {
		Val     string
		Left    *Node
		Right   *Node
		Parent  *Node
		Balance int
	}
	node1 := &Node{
		Val:     "(",
		Balance: 1,
	}
	type funcType func(root *Node, i int)
	var fun funcType
	fun = func(node *Node, i int) {
		if i == n-1 {
			return
		}
		if node.Balance > 0 {
			node.Left = &Node{
				Val:     "(",
				Balance: node.Balance + 1,
				Parent:  node,
			}
			node.Right = &Node{
				Val:     ")",
				Balance: node.Balance - 1,
				Parent:  node,
			}
			fun(node.Left, i+1)
			fun(node.Right, i+1)
		} else {
			node.Left = &Node{
				Val:     "(",
				Balance: node.Balance + 1,
				Parent:  node,
			}
			fun(node.Left, i+1)
		}
	}
	fun(node1, 0)
	node2 := &Node{
		Val:     ")",
		Balance: -1,
	}
	fun = func(node *Node, i int) {
		if i == n-1 {
			return
		}
		if node.Balance < 0 {
			node.Left = &Node{
				Val:     "(",
				Balance: node.Balance + 1,
				Parent:  node,
			}
			node.Right = &Node{
				Val:     ")",
				Balance: node.Balance - 1,
				Parent:  node,
			}
			fun(node.Left, i+1)
			fun(node.Right, i+1)
		} else {
			node.Right = &Node{
				Val:     ")",
				Balance: node.Balance - 1,
				Parent:  node,
			}
			fun(node.Right, i+1)
		}
	}
	fun(node2, 0)
	dict := make(map[string]int)
	elem := ""
	var curr *Node
	type preOrder func(root *Node)
	var po1 preOrder
	po1 = func(node *Node) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			curr = node
			for {
				elem = node.Val + elem
				node = node.Parent
				if node.Parent == nil {
					elem = node.Val + elem
					break
				}
			}
			node = curr
			dict[elem] = node.Balance
			elem = ""
		}
		po1(node.Left)
		po1(node.Right)
	}
	po1(node1)
	var po2 preOrder
	po2 = func(node *Node) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			curr = node
			for {
				elem = elem + node.Val
				node = node.Parent
				if node.Parent == nil {
					elem = elem + node.Val
					break
				}
			}
			node = curr
			dict[elem] = node.Balance
			elem = ""
		}
		po2(node.Left)
		po2(node.Right)
	}
	po2(node2)
	for k1, v1 := range dict {
		if v1 >= 0 {
			for k2, v2 := range dict {
				if v1+v2 == 0 {
					arr = append(arr, k1+k2)
				}
			}
		}
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	arr := generateParenthesis(2)
//	for _, elem := range arr {
//		fmt.Println(elem)
//	}
//}

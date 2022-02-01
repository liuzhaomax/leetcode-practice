/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-27 22:03:21
 * @version     v1.0
 * @filename    go.[32]最长有效括号
 * @description
 ***************************************************************************/
package main

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
// Related Topics 栈 字符串 动态规划 👍 1626 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	res := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && (i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' && s[i] == ')') {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		res = max(res, dp[i])
	}
	return res
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	s := "((())"
//	fmt.Println(longestValidParentheses(s))
//}

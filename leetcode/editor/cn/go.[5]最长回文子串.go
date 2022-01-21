/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-21 20:39:18
 * @version     v1.0
 * @filename    go.[5]最长回文子串
 * @description
 ***************************************************************************/
package main

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划 👍 4602 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	length := len(s)
	if length == 1 {
		return s
	}
	max := ""
	for i := 0; i < length; i++ {
		for j := 1; j < length; j++ {
			if i-j < 0 || i+j >= length {
				break
			}
			if s[i-j] != s[i+j] {
				break
			}
			if len(max) < 2*j+1 {
				max = s[i-j : i+j+1]
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := 1; j < length; j++ {
			if i-j+1 < 0 || i+j >= length {
				break
			}
			if s[i-j+1] != s[i+j] {
				break
			}
			if len(max) < 2*j {
				max = s[i-j+1 : i+j+1]
			}
		}
	}
	if max == "" {
		return string(s[0])
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

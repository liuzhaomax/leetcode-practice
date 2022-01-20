/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-21 00:49:46
 * @version     v1.0
 * @filename    go.[3]无重复字符的最长子串
 * @description
 ***************************************************************************/
package main

import "strings"

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口 👍 6788 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	str := ""
	for i := 0; i < len(s); i++ {
		if strings.Contains(str, string(s[i])) {
			for j := 0; j < len(str); j++ {
				if str[j] == s[i] {
					str = str[j+1:]
				}
			}
		}
		str = str + string(s[i])
		if max < len(str) {
			max = len(str)
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

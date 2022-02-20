/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 02:47:52
 * @version     v1.0
 * @filename    go.[131]分割回文串
 * @description
 ***************************************************************************/
package main

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
// Related Topics 字符串 动态规划 回溯 👍 988 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}
	var res [][]string
	partitionBT(&res, 0, make([]string, 0), s)
	return res
}

func partitionBT(resPtr *[][]string, start int, current []string, s string) {
	if start == len(s) {
		tmp := make([]string, len(current))
		copy(tmp, current)
		*resPtr = append(*resPtr, tmp)
	} else {
		for i := start; i < len(s); i++ {
			if partitionCheck(s, start, i) {
				current = append(current, s[start:i+1])
				partitionBT(resPtr, i+1, current, s)
				current = current[:len(current)-1]
			}
		}
	}
}

func partitionCheck(s string, left int, right int) bool {
	for left < right {
		if s[left] == s[right] {
			left += 1
			right -= 1
		} else {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	s := "aab"
//	fmt.Println(partition(s))
//}

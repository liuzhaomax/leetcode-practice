/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-03 01:08:46
 * @version     v1.0
 * @filename    go.[76]最小覆盖子串
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
)

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？ Related Topics 哈希表 字符串 滑动窗口 👍 1612 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	// corner case
	if len(t) == 0 {
		return ""
	}
	var (
		countT                      = make(map[byte]int) // count the number of occurrences of characters in t
		need                    int = 0                  // the number of requirements needs to be matched
		window                      = make(map[byte]int) // count the number of occurrences of characters in the window
		have                    int = 0                  // the number of current requirements that the window is matching
		left, right             int = 0, 0               // window pointers
		resultLeft, resultRight int = 0, 0               // result window
		resultLen               int = 100000
	)

	// build maps
	for i := range t {
		// window - avoid panic
		window[t[i]] = 0
		// countT map
		countT[t[i]]++
	}
	// init the number of requirements
	need = len(countT)

	for i := range s {
		// expand the window
		right = i

		// check if the current character is needed to be considered
		_, ok := countT[s[i]]
		if !ok {
			continue
		}

		// increase the window count map
		window[s[i]] += 1
		// check requirements
		if window[s[i]] == countT[s[i]] {
			have++
		}

		// the current window is having all of the characters in t
		// try to shrink the window by removing the leftmost character
		for have == need {
			if (right - left + 1) < resultLen {
				resultLeft = left
				resultRight = right
				resultLen = (right - left + 1)
			}

			// remove a character from the window
			_, ok := countT[s[left]]
			if !ok { // not worth to consider
				left += 1 // shrink the window
				continue
			}

			// decrease the window count
			window[s[left]] -= 1

			// update the number of matching requirements
			if window[s[left]] < countT[s[left]] {
				have -= 1
			}
			left += 1
		}
	}

	// find nothing
	if resultLen == 100000 {
		return ""
	}

	return s[resultLeft : resultRight+1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s := "ababb"
	t := "abb"
	fmt.Println(minWindow(s, t))
	//fmt.Println(strings.Contains(s, t))
}

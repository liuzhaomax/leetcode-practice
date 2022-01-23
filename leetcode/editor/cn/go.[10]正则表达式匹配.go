/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-21 21:15:58
 * @version     v1.0
 * @filename    go.[10]正则表达式匹配
 * @description
 ***************************************************************************/
package main

import "fmt"

//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
//
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//
// 示例 1：
//
//
//输入：s = "aa" p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。
//
//
// 示例 2:
//
//
//输入：s = "aa" p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
//
// 示例 3：
//
//
//输入：s = "ab" p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
//
// 示例 4：
//
//
//输入：s = "aab" p = "c*a*b"
//输出：true
//解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
//
//
// 示例 5：
//
//
//输入：s = "mississippi" p = "mis*is*p*."
//输出：false
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// 1 <= p.length <= 30
// s 只含小写英文字母。
// p 只含小写英文字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
//
// Related Topics 递归 字符串 动态规划 👍 2669 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	lens := len(s) + 1
	lenp := len(p) + 1
	dp := make([][]bool, lenp)
	for i := range dp {
		dp[i] = make([]bool, lens)
	}
	dp[lenp-1][lens-1] = true
	star := false
	flag := false
	for i := lenp - 2; i >= 0; i-- {
		if star == true {
			for k := 0; k < lens; k++ {
				flag = flag || dp[i][k]
			}
			if flag == false {
				for k := 0; k < lens; k++ {
					dp[i][k] = true
				}
			}
			flag = false
			star = false
		}
		for j := lens - 2; j >= 0; j-- {
			switch string(p[i]) {
			case "*":
				star = true
				if dp[i+1][j+1] == true {
					dp[i][j] = true
				}
				if star == true && s[j] == p[i-1] && dp[i][j+1] == true {
					dp[i-1][j] = true
				}
			case ".":
				dp[i][j] = true
			default:
				if s[j] == p[i] && dp[i+1][j+1] == true {
					dp[i][j] = true
				}

			}
		}
		if star == false {
			flag = false
			for k := 0; k < lens; k++ {
				flag = flag || dp[i][k]
			}
			if flag == false {
				continue
			}
		}
	}
	return dp[0][0]
}

//func isMatch(s string, p string) bool {
//	lens := len(s)
//	lenp := len(p)
//	s = " " + s
//	p = " " + p
//	dp := make([][]bool, lenp+1)
//	for i := range dp {
//		dp[i] = make([]bool, lens+1)
//	}
//	dp[0][0] = true
//	for i := 0; i <= lens; i++ {
//		for j := 1; j <= lenp; j++ {
//			if j+1 <= lenp && string(p[j+1]) == "*" {
//				continue
//			}
//			if string(p[j]) != "*" {
//				dp[i][j] = i > 0 && j > 0 && dp[i-1][j-1] && (s[i] == p[j] || string(p[j]) == ".")
//			} else {
//				dp[i][j] = j >= 2 && dp[i][j-2] || (i > 0 && j > 0 && dp[i-1][j] && (s[i] == p[j-1] || string(p[j-1]) == "."))
//			}
//		}
//	}
//	return dp[lens][lenp]
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	//var aaa = "aab"
	//var bbb = "c*a*b"
	//fmt.Println(isMatch(aaa, bbb)) //true
	//var ggg = "aaa"
	//var hhh = "ab*ac*a"
	//fmt.Println(isMatch(ggg, hhh)) //true
	//var ccc = "miss"
	//var ddd = "m"
	//fmt.Println(isMatch(ccc, ddd)) //false
	//var eee = "abcd"
	//var fff = "d*"
	//fmt.Println(isMatch(eee, fff)) //false
	//var iii = "aa"
	//var jjj = "a"
	//fmt.Println(isMatch(iii, jjj)) //false
	//var kkk = "mississippi"
	//var lll = "mis*is*ip*."
	//fmt.Println(isMatch(kkk, lll)) //true
	var mmm = "mississippi"
	var nnn = "mis*is*p*."
	fmt.Println(isMatch(mmm, nnn)) //false
}

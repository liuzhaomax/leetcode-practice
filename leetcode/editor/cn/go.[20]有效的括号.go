/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-25 15:38:00
 * @version     v1.0
 * @filename    go.[20]有效的括号
 * @description
 ***************************************************************************/
package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串 👍 2925 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	arr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(":
			arr = append(arr, ")")
		case "{":
			arr = append(arr, "}")
		case "[":
			arr = append(arr, "]")
		case ")":
			if len(arr) > 0 && arr[len(arr)-1] == ")" {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		case "}":
			if len(arr) > 0 && arr[len(arr)-1] == "}" {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		case "]":
			if len(arr) > 0 && arr[len(arr)-1] == "]" {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}
	if len(arr) == 0 {
		return true
	} else {
		return false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	a := "()"
//	fmt.Println(isValid(a))
//}

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 22:43:27
 * @version     v1.0
 * @filename    go.[139]单词拆分
 * @description
 ***************************************************************************/
package main

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
//
//
// 示例 1：
//
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
//
// 示例 2：
//
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//     注意，你可以重复使用字典中的单词。
//
//
// 示例 3：
//
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中的所有字符串 互不相同
//
// Related Topics 字典树 记忆化搜索 哈希表 字符串 动态规划 👍 1404 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return false
	}
	var jump int
	dp := make([]bool, len(wordDict)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if i < jump {
			continue
		}
		for j := 0; j < len(wordDict); j++ {
			if i+len(wordDict[j]) <= len(s) && s[i:i+len(wordDict[j])] == wordDict[j] {
				dp[j+1] = dp[j]
				jump = i + len(wordDict[j])
				break
			}
		}
	}
	return dp[len(wordDict)]
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	//s := "leetcode"
//	//wordDict := []string{"leet", "code"}
//	s := "a"
//	wordDict := []string{"b"}
//	fmt.Println(wordBreak(s, wordDict))
//}

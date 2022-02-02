/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-02 19:46:29
 * @version     v1.0
 * @filename    go.[49]字母异位词分组
 * @description
 ***************************************************************************/
package main

import (
	"sort"
	"strings"
)

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
//
//
//
// 示例 1:
//
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
//
//
//输入: strs = [""]
//输出: [[""]]
//
//
// 示例 3:
//
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
// Related Topics 哈希表 字符串 排序 👍 983 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	strsCopied := make([]string, len(strs))
	copy(strsCopied, strs)
	res := [][]string{}
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		strs[i] = sortString(strs[i])
	}
	for i := 0; i < len(strs); i++ {
		m[strs[i]] = append(m[strs[i]], strsCopied[i])
	}
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//	fmt.Println(groupAnagrams(strs))
//}

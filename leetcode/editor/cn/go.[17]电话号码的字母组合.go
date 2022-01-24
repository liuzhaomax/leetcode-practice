/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-24 22:36:28
 * @version     v1.0
 * @filename    go.[17]电话号码的字母组合
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"math"
	"strconv"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 哈希表 字符串 回溯 👍 1689 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	var arr []string
	if digits == "" {
		return arr
	}
	m := make(map[int]string)
	m[2] = "abc"
	m[3] = "def"
	m[4] = "ghi"
	m[5] = "jkl"
	m[6] = "mno"
	m[7] = "pqrs"
	m[8] = "tuv"
	m[9] = "wxyz"
	length := 1
	for _, num := range digits {
		if string(num) == "7" || string(num) == "9" {
			length = length * 4
		} else {
			length = length * 3
		}
	}
	for i := 0; i < length; i++ {
		arr = append(arr, "")
	}
	product := 1
	for i := len(digits) - 1; i >= 0; i-- {
		for k, v := range m {
			if strconv.Itoa(k) == string(digits[i]) {
				for j := 0; j < length; j++ {
					arr[j] = string(v[int(math.Floor(float64(j/product)))%len(v)]) + arr[j]
				}
				product = product * len(v)
				break
			}
		}
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	a := "7"
	fmt.Println(letterCombinations(a))
	//d := "23"
	//fmt.Println(letterCombinations(d))
}

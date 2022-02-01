/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-02 00:27:40
 * @version     v1.0
 * @filename    go.[46]全排列
 * @description
 ***************************************************************************/
package main

import "fmt"

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// Related Topics 数组 回溯 👍 1750 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	current := []int{}
	visited := make([]bool, len(nums))
	permuteBT(nums, &res, current, visited)
	return res
}

func permuteBT(nums []int, resPtr *[][]int, current []int, visited []bool) {
	if len(nums) == len(current) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*resPtr = append(*resPtr, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}
		current = append(current, nums[i])
		visited[i] = true
		permuteBT(nums, resPtr, current, visited)
		visited[i] = false
		current = current[:len(current)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permute(nums))
}

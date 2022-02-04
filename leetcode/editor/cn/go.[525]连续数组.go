/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-04 23:00:13
 * @version     v1.0
 * @filename    go.[525]连续数组
 * @description
 ***************************************************************************/
package main

import "fmt"

//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
//
// 示例 2:
//
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 不是 0 就是 1
//
// Related Topics 数组 哈希表 前缀和 👍 514 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func findMaxLength(nums []int) int {
//	if len(nums) == 0 || len(nums) == 1 {
//		return 0
//	}
//	max := 0
//	for j := 0; j < len(nums); j++ {
//		if nums[j] == 0 {
//			nums[j] = -1
//		}
//		count := nums[j]
//		start := j
//		for i := j + 1; i < len(nums); i++ {
//			if nums[i] == 0 {
//				nums[i] = -1
//			}
//			count += nums[i]
//			if count == 0 {
//				max = findMaxLengthMax(max, i-start+1)
//			}
//		}
//	}
//	return max
//}
//func findMaxLengthMax(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func findMaxLength(nums []int) int {
	// valIdx stores the first index of a sum
	dp := make(map[int]int, len(nums)/2)
	dp[0] = -1
	var res int
	var count int
	for i, n := range nums {
		if n == 0 {
			n = -1
		}
		count += n
		if idx, exists := dp[count]; exists {
			if d := i - idx; d > res {
				res = d
			}
		} else {
			dp[count] = i
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	arr := []int{1, 1, 1, 0, 1, 0}
	fmt.Println(findMaxLength(arr))
}

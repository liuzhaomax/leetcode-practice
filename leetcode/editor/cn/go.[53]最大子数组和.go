/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-02 20:35:19
 * @version     v1.0
 * @filename    go.[53]最大子数组和
 * @description
 ***************************************************************************/
package main

import (
	"math"
)

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：1
//
//
// 示例 3：
//
//
//输入：nums = [5,4,-1,7,8]
//输出：23
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
// Related Topics 数组 分治 动态规划 👍 4296 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func maxSubArray(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	sum := 0
//	tmp := 0
//	max := nums[0]
//	for i := 0; i < len(nums); i++ {
//		tmp = sum + nums[i]
//		if tmp >= nums[i] {
//			sum = tmp
//		} else {
//			sum = nums[i]
//		}
//		max = int(math.Max(float64(max), float64(sum)))
//	}
//	return max
//}

func maxSubArray(nums []int) int {
	var localmax, globalmax int
	localmax = nums[0]
	globalmax = nums[0]
	//kadane's algorithm approach
	for i := 1; i < len(nums); i++ {
		localmax = int(math.Max(float64(localmax+nums[i]), float64(nums[i])))
		globalmax = int(math.Max(float64(localmax), float64(globalmax)))
	}
	return globalmax
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{-2, -1}
//	fmt.Println(maxSubArray(nums))
//}

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-06 15:30:07
 * @version     v1.0
 * @filename    go.[152]乘积最大子数组
 * @description
 ***************************************************************************/
package main

//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
// 示例 1:
//
// 输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
// 输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
// Related Topics 数组 动态规划 👍 1476 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dpMin := make([]int, len(nums))
	dpMin[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxProductMax(nums[i], dp[i-1]*nums[i], dpMin[i-1]*nums[i])
		dpMin[i] = maxProductMin(nums[i], dp[i-1]*nums[i], dpMin[i-1]*nums[i])
		max = maxProductMax2(max, dp[i])
	}
	return max
}

func maxProductMax2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProductMax(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func maxProductMin(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{-1, -1, -2, -2}
//	fmt.Println(maxProduct(nums))
//}

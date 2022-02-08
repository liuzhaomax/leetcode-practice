/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-06 21:17:42
 * @version     v1.0
 * @filename    go.[213]打家劫舍 II
 * @description
 ***************************************************************************/
package main

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 3：
//
//
//输入：nums = [1,2,3]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
// Related Topics 数组 动态规划 👍 911 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return robMax2(nums[0], nums[1])
	}
	dp := make([]int, len(nums)-1)
	dp[0] = nums[1]
	dp[1] = robMax2(nums[1], nums[2])
	dp1 := make([]int, len(nums)-1)
	dp1[0] = nums[0]
	dp1[1] = robMax2(nums[0], nums[1])
	return robMax2(robb(nums, dp), robb2(nums, dp1))
}

func robb(nums []int, dp []int) int {
	for i := 2; i < len(dp); i++ {
		dp[i] = robMax2(dp[i-1], dp[i-2]+nums[i+1])
	}
	return dp[len(dp)-1]
}

func robb2(nums []int, dp []int) int {
	for i := 2; i < len(dp); i++ {
		dp[i] = robMax2(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}

func robMax2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	nums := []int{1, 2, 3, 1}
//	fmt.Println(rob(nums))
//}

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-01 23:21:06
 * @version     v1.0
 * @filename    go.[45]跳跃游戏 II
 * @description
 ***************************************************************************/
package main

//给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 假设你总是可以到达数组的最后一个位置。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 示例 2:
//
//
//输入: nums = [2,3,0,1,4]
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
//
// Related Topics 贪心 数组 动态规划 👍 1388 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func jump(nums []int) int {
//	dp := []int{}
//	target := nums[len(nums)-1]
//	for i := 0; i < len(nums) - 1; i++ {
//		if i + nums[i] == target {
//			dp = append(dp, i)
//		}
//	}
//	start := 1
//}
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if dp[len(dp)-1] > 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if dp[j] != 0 {
				continue
			}
			if j-i > nums[i] {
				break
			}
			if j-i <= nums[i] && dp[j] == 0 {
				dp[j] = dp[i] + 1
			}
		}
	}
	return dp[len(dp)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	nums := []int{2, 3, 0, 1, 4}
//	fmt.Println(jump(nums))
//}

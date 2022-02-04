/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-05 03:25:08
 * @version     v1.0
 * @filename    go.[416]分割等和子集
 * @description
 ***************************************************************************/
package main

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
// Related Topics 数组 动态规划 👍 1122 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	for i := range dp {
		dp[i][0] = true
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < sum+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}

//leetcode submit region end(Prohibit modification and deletion)
//[true true false false false false]
//[true true true true false false]
//[true true true true true true]
//[true true true true true true]

//func main() {
//	nums := []int{1, 2, 3, 5}
//	fmt.Println(canPartition(nums))
//}

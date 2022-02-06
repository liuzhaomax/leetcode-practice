/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-07 01:27:20
 * @version     v1.0
 * @filename    go.[123]买卖股票的最佳时机 III
 * @description
 ***************************************************************************/
package main

//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入：prices = [3,3,5,0,0,3,1,4]
//输出：6
//解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//
// 示例 2：
//
//
//输入：prices = [1,2,3,4,5]
//输出：4
//解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
//
// 示例 3：
//
//
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
//
// 示例 4：
//
//
//输入：prices = [1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 10⁵
// 0 <= prices[i] <= 10⁵
//
// Related Topics 数组 动态规划 👍 998 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit3(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	tmpMax := 0
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, len(prices))
	}
	for i := 1; i < 3; i++ {
		tmpMax = -prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = maxProfit3Max(dp[i][j-1], prices[j]+tmpMax)
			tmpMax = maxProfit3Max(tmpMax, dp[i-1][j]-prices[j])
		}
	}
	return dp[2][len(prices)-1]
}

func maxProfit3Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
//	//prices := []int{1, 2, 3, 4, 5}
//	fmt.Println(maxProfit3(prices))
//}

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-02 22:20:05
 * @version     v1.0
 * @filename    go.[64]最小路径和
 * @description
 ***************************************************************************/
package main

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// Related Topics 数组 动态规划 矩阵 👍 1129 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 && i > 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if j > 0 && i > 0 {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

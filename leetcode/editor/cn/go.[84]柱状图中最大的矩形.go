/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 00:15:27
 * @version     v1.0
 * @filename    go.[84]柱状图中最大的矩形
 * @description
 ***************************************************************************/
package main

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
// 示例 1:
//
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//
// 示例 2：
//
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//
// 提示：
//
//
// 1 <= heights.length <=10⁵
// 0 <= heights[i] <= 10⁴
//
// Related Topics 栈 数组 单调栈 👍 1754 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	start := 0
	end := len(heights) - 1
	max := heights[0]
	for i := 0; i < len(heights); i++ {
		if i > 0 && heights[i] == heights[i-1] {
			continue
		}
		start = 0
		end = len(heights) - 1
		for j := i; j >= 0; j-- {
			if j < i && heights[j] < heights[i] {
				start = j + 1
				break
			}
		}
		for j := 0; j < len(heights); j++ {
			if j > i && heights[j] < heights[i] {
				end = j - 1
				break
			}
		}
		max = largestRectangleAreaMax(max, heights[i]*(end-start+1))
	}
	return max
}

func largestRectangleAreaMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	heights := []int{4, 2, 4, 5, 3}
//	fmt.Println(largestRectangleArea(heights))
//}

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-01 16:26:13
 * @version     v1.0
 * @filename    go.[42]接雨水
 * @description
 ***************************************************************************/
package main

import (
	"math"
)

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 3060 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func trap(height []int) int {
//	if len(height) == 0 || len(height) == 1 {
//		return 0
//	}
//	res := 0
//	tmp := 0
//	left := 0
//	right := 0
//	min := height[0]
//	max := height[0]
//	for i := 0; i < len(height); i++ {
//		if height[i] < min {
//			min = height[i]
//		}
//		if height[i] > max {
//			max = height[i]
//		}
//	}
//	for i := min; i <= max; i++ {
//		for j := 0; j < len(height); j++ {
//			if height[j] >= i {
//				right = j
//			}
//			if j > 0 && height[j-1] < i && right > left && right >= 2 && height[left] >= i {
//				tmp += right - left - 1
//			}
//			if height[j] >= i {
//				left = j
//			}
//		}
//		res += tmp
//		tmp = 0
//		left = 0
//		right = 0
//	}
//	return res
//}

//func trap(height []int) int {
//	if len(height) == 0 || len(height) == 1 {
//		return 0
//	}
//	res := 0
//	tmp := 0
//	left := 0
//	right := len(height) - 1
//	min := height[0]
//	max := height[0]
//	for i := 0; i < len(height); i++ {
//		if height[i] < min {
//			min = height[i]
//		}
//		if height[i] > max {
//			max = height[i]
//		}
//	}
//	for i := min; i <= max; i++ {
//		for left < right-1 {
//			if height[left] < i {
//				left++
//			}
//			if height[right] < i {
//				right--
//			}
//			if height[left] >= i && height[right] >= i {
//				break
//			}
//		}
//		if left > right-2 {
//			break
//		}
//		for j := left; j < right; j++ {
//			if height[j] < i {
//				tmp += 1
//			}
//		}
//		res += tmp
//		tmp = 0
//	}
//	return res
//}

func trap(height []int) int {
	_, maxIndex := maxIntSlice(height)
	rain := 0

	// loop from left to max
	maxHeight := 0
	for i := 0; i < maxIndex; i++ {
		maxHeight = maxInt(maxHeight, height[i])
		rain += maxHeight - height[i]
	}

	// loop from right to max
	maxHeight = 0
	for i := len(height) - 1; i > maxIndex; i-- {
		maxHeight = maxInt(maxHeight, height[i])
		rain += maxHeight - height[i]
	}

	return rain
}
func maxIntSlice(nums []int) (max, maxIndex int) {
	max = math.MinInt64
	maxIndex = -1
	for i, n := range nums {
		if n > max {
			max = n
			maxIndex = i
		}
	}
	return
}
func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//	height := []int{4, 2, 0, 3, 2, 5}
//	//height := []int{9, 6, 8, 8, 5, 6, 3}
//	fmt.Println(trap(height))
//}

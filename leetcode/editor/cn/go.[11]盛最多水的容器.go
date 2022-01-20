/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-20 20:47:33
 * @version     v1.0
 * @filename    go.[11]盛最多水的容器
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"math"
)

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
// 示例 3：
//
//
//输入：height = [4,3,2,1,4]
//输出：16
//
//
// 示例 4：
//
//
//输入：height = [1,2,1]
//输出：2
//
//
//
//
// 提示：
//
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
//
// Related Topics 贪心 数组 双指针 👍 3126 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	maxProduct := int(math.Min(float64(height[0]), float64(height[len(height)-1])) * float64(len(height)-1))
	for i := 0; i < len(height)-1; i++ {
		if i > 0 {
			if height[i-1] >= height[i] {
				continue
			}
		}
		for j := len(height) - 1; j > i; j-- {
			maxProduct = int(math.Max(math.Min(float64(height[i]), float64(height[j]))*float64(j-i), float64(maxProduct)))
			if j < len(height)-1 {
				if height[j] >= height[i] {
					break
				}
			}
		}
	}
	return maxProduct
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	arr := []int{100, 0, 0, 0, 0, 100}
	fmt.Println(maxArea(arr))
}

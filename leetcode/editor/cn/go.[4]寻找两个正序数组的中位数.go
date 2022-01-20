/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021-11-12 16:34:23
 * @version     v1.0
 * @filename    go.[4]寻找两个正序数组的中位数
 * @description
 ***************************************************************************/
package main

import (
	"math"
	"sort"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
// 示例 3：
//
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//
//
// 示例 4：
//
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//
//
// 示例 5：
//
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
// Related Topics 数组 二分查找 分治 👍 4649 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	length := len(nums)
	var pos int
	if length%2 == 0 {
		pos = int(math.Ceil(float64(length / 2)))
	} else {
		pos = length / 2
	}
	return float64(nums[pos])
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	nums1 := []int{1,3}
//	nums2 := []int{2}
//	result := findMedianSortedArrays(nums1, nums2)
//	fmt.Println(result)
//}

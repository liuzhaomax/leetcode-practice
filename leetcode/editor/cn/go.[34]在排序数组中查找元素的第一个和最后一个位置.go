/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-28 23:36:09
 * @version     v1.0
 * @filename    go.[34]在排序数组中查找元素的第一个和最后一个位置
 * @description
 ***************************************************************************/
package main

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 进阶：
//
//
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
// 示例 2：
//
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
// 示例 3：
//
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 二分查找 👍 1425 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := 0
	right := len(nums) - 1
	l := -1
	r := -1
	prevLeft := -1
	prevRight := -1
	for left <= right {
		if prevLeft == left && prevRight == right {
			break
		}
		prevLeft = left
		prevRight = right
		if nums[left] != target {
			left++
		} else {
			l = left
		}
		if nums[right] != target {
			right--
		} else {
			r = right
		}
	}
	return []int{l, r}
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{5, 7, 7, 8, 8, 10}
//	fmt.Println(searchRange(nums, 10))
//}

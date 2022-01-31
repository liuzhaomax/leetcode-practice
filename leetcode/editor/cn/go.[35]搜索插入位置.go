/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-31 18:06:49
 * @version     v1.0
 * @filename    go.[35]搜索插入位置
 * @description
 ***************************************************************************/
package main

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//
//
// 示例 3:
//
//
//输入: nums = [1,3,5,6], target = 7
//输出: 4
//
//
// 示例 4:
//
//
//输入: nums = [1,3,5,6], target = 0
//输出: 0
//
//
// 示例 5:
//
//
//输入: nums = [1], target = 0
//输出: 0
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 为无重复元素的升序排列数组
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 二分查找 👍 1321 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func searchInsert(nums []int, target int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] >= target {
//			return i
//		}
//	}
//	return len(nums)
//}
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)

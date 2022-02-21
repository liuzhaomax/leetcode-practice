/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-21 23:36:06
 * @version     v1.0
 * @filename    go.[283]移动零
 * @description
 ***************************************************************************/
package main

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
//输入: nums = [0]
//输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
// Related Topics 数组 双指针 👍 1442 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] != 0 {
			i++
			j++
			continue
		}
		if nums[i] == 0 && nums[j] == 0 {
			j++
			continue
		}
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{0, 1, 0, 3, 12}
//	moveZeroes(nums)
//	fmt.Println(nums)
//}

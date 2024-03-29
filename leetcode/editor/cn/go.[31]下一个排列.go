/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-26 19:17:59
 * @version     v1.0
 * @filename    go.[31]下一个排列
 * @description
 ***************************************************************************/
package main

import (
	"sort"
)

//整数数组的一个 排列 就是将其所有成员以序列或线性顺序排列。
//
//
// 例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
//
//
// 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就
//是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
//
//
// 例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
// 类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
// 而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
//
//
// 给你一个整数数组 nums ，找出 nums 的下一个排列。
//
// 必须 原地 修改，只允许使用额外常数空间。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//
//
// 示例 3：
//
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
//
// Related Topics 数组 双指针 👍 1501 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	if len(nums) == 2 {
		nums[0] = nums[0] + nums[1]
		nums[1] = nums[0] - nums[1]
		nums[0] = nums[0] - nums[1]
		return
	}
	decrement := true
	index := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 && nums[i] < nums[i+1] {
			decrement = false
			index = i
			break
		}
	}
	if decrement == true {
		sort.Ints(nums)
		return
	} else {
		if index == len(nums)-2 {
			nums[len(nums)-2] = nums[len(nums)-2] + nums[len(nums)-1]
			nums[len(nums)-1] = nums[len(nums)-2] - nums[len(nums)-1]
			nums[len(nums)-2] = nums[len(nums)-2] - nums[len(nums)-1]
			return
		}
		flag := false
		for i := index; i < len(nums)-1; i++ {
			if nums[i+1] <= nums[index] {
				nums[index] = nums[index] + nums[i]
				nums[i] = nums[index] - nums[i]
				nums[index] = nums[index] - nums[i]
				flag = true
				break
			}
		}
		if flag == false {
			nums[index] = nums[index] + nums[len(nums)-1]
			nums[len(nums)-1] = nums[index] - nums[len(nums)-1]
			nums[index] = nums[index] - nums[len(nums)-1]
		}
		sort.Ints(nums[index+1:])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{1, 3, 2}
//	nextPermutation(nums)
//	fmt.Println(nums)
//}

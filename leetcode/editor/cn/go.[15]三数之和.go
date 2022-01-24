/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-24 21:01:51
 * @version     v1.0
 * @filename    go.[15]三数之和
 * @description
 ***************************************************************************/
package main

import (
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4232 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	arr := make([]int, 3)
	var array [3]int
	mapp := make(map[[3]int]int)
	for k := 0; k < len(nums)-2; k++ {
		for i := k + 1; i < len(nums)-1; i++ {
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			for j := len(nums) - 1; j > i; j-- {
				if j < len(nums)-2 && nums[j] == nums[j+1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] < 0 {
					break
				} else if nums[i]+nums[j]+nums[k] > 0 {
					continue
				} else {
					array[0] = nums[k]
					array[1] = nums[i]
					array[2] = nums[j]
					mapp[array] = 1
				}
			}
		}
	}
	for key := range mapp {
		for i, val := range key {
			arr[i] = val
		}
		result = append(result, arr)
		arr = make([]int, 3)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	arr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
//	fmt.Println(threeSum(arr))
//}

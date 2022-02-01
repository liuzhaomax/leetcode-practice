/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-01 15:51:48
 * @version     v1.0
 * @filename    go.[41]缺失的第一个正数
 * @description
 ***************************************************************************/
package main

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
// Related Topics 数组 哈希表 👍 1329 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 && nums[0] <= 0 {
		return 1
	}
	dict := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] = true
	}
	for i := 1; i <= len(nums)+1; i++ {
		if dict[i] == false {
			return i
		}
	}
	return 1
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	arr := []int{1, 2, 0}
//	fmt.Println(firstMissingPositive(arr))
//}

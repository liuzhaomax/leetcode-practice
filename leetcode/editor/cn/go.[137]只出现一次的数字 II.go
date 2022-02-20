/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 22:36:58
 * @version     v1.0
 * @filename    go.[137]只出现一次的数字 II
 * @description
 ***************************************************************************/
package main

//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,3,2]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,1,0,1,99]
//输出：99
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
// Related Topics 位运算 数组 👍 806 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var m = make(map[int]int)
	for i := range nums {
		m[nums[i]] += 1
	}
	var res int
	for k, val := range m {
		if val == 1 {
			res = k
			break
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{2, 2, 3, 2}
//	fmt.Println(singleNumber2(nums))
//}

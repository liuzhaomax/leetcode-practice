/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 22:20:38
 * @version     v1.0
 * @filename    go.[136]只出现一次的数字
 * @description
 ***************************************************************************/
package main

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 说明：
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
// 示例 1:
//
// 输入: [2,2,1]
//输出: 1
//
//
// 示例 2:
//
// 输入: [4,1,2,1,2]
//输出: 4
// Related Topics 位运算 数组 👍 2253 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
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
//	nums := []int{4, 1, 2, 2, 1}
//	fmt.Println(singleNumber(nums))
//}

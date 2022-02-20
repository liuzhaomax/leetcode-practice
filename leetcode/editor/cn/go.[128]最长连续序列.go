/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-20 21:53:59
 * @version     v1.0
 * @filename    go.[128]最长连续序列
 * @description
 ***************************************************************************/
package main

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
// Related Topics 并查集 数组 哈希表 👍 1101 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	max := 0
	for num, _ := range m {
		if !m[num-1] {
			start := num
			cnt := 0
			for m[start] {
				cnt++
				start++
			}
			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	nums := []int{100, 4, 200, 3, 1, 2}
//	fmt.Println(longestConsecutive(nums))
//}

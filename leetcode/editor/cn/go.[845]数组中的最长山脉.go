/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-02 15:52:03
 * @version     v1.0
 * @filename    go.[845]数组中的最长山脉
 * @description
 ***************************************************************************/
package main

//把符合下列属性的数组 arr 称为 山脉数组 ：
//
//
// arr.length >= 3
// 存在下标 i（0 < i < arr.length - 1），满足
//
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//
//
//
//
// 给出一个整数数组 arr，返回最长山脉子数组的长度。如果不存在山脉子数组，返回 0 。
//
//
//
// 示例 1：
//
//
//输入：arr = [2,1,4,7,3,2,5]
//输出：5
//解释：最长的山脉子数组是 [1,4,7,3,2]，长度为 5。
//
//
// 示例 2：
//
//
//输入：arr = [2,2,2]
//输出：0
//解释：不存在山脉子数组。
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10⁴
// 0 <= arr[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你可以仅用一趟扫描解决此问题吗？
// 你可以用 O(1) 空间解决此问题吗？
//
// Related Topics 数组 双指针 动态规划 枚举 👍 220 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func longestMountain(arr []int) int {
//	res := 0
//	left := -1
//	start := 0
//	tmp := 0
//	for i := 0; i < len(arr); i++ {
//		if (i > 0 && arr[i-1] >= arr[i] || i == 0) && (i < len(arr)-1 && arr[i] < arr[i+1] || i == len(arr)-1) {
//			start += 1
//		}
//		if i > 0 && arr[i-1] > arr[i] && (i < len(arr)-1 && arr[i] <= arr[i+1] || i == len(arr)-1) && left >= 0 && start == 2 {
//			tmp += i - left + 1
//			left = -1
//			start = 0
//		}
//		if (i > 0 && arr[i-1] >= arr[i] || i == 0) && (i < len(arr)-1 && arr[i] < arr[i+1] || i == len(arr)-1) {
//			left = i
//			res = max(res, tmp)
//			tmp = 0
//		}
//	}
//	return res
//}
//
//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	//arr := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
//	arr := []int{0, 2, 0, 2, 1, 2, 3, 4, 4, 1}
//	fmt.Println(longestMountain(arr))
//}

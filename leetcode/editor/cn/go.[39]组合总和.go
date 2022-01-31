/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-01-31 18:21:27
 * @version     v1.0
 * @filename    go.[39]组合总和
 * @description
 ***************************************************************************/
package main

import "fmt"

//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的
// 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
//
//
// 示例 1：
//
//
//输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。
//
// 示例 2：
//
//
//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]
//
// 示例 3：
//
//
//输入: candidates = [2], target = 1
//输出: []
//
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都 互不相同
// 1 <= target <= 500
//
// Related Topics 数组 回溯 👍 1734 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	if len(candidates) == 0 {
		return res
	}
	combinationSumSearch(&res, cur, candidates, target, 0)
	return res
}

func combinationSumSearch(res *[][]int, cur []int, candidates []int, target int, level int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := level; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		combinationSumSearch(res, cur, candidates, target-candidates[i], i)
		cur = cur[:len(cur)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	candidates := []int{7, 3, 9, 6}
	target := 6
	fmt.Println(combinationSum(candidates, target))
}

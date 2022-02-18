/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-19 00:23:07
 * @version     v1.0
 * @filename    go.[79]单词搜索
 * @description
 ***************************************************************************/
package main

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"SEE"
//输出：true
//
//
// 示例 3：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCB"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
// Related Topics 数组 回溯 矩阵 👍 1186 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	var (
		rows = len(board)
		cols = len(board[0])
		path = make([][]bool, rows)
	)
	for i := range path {
		path[i] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if existBT(board, row, col, word, 0, &path) == true {
				return true
			}
		}
	}

	return false
}

func existBT(board [][]byte, row, col int, word string, wordIDX int, path *[][]bool) bool {
	// found the last character - (in the previous)
	if wordIDX == len(word) {
		return true
	}
	// out of bound
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
		return false
	}
	// not match
	if board[row][col] != word[wordIDX] {
		return false
	}
	// in the existing path
	if (*path)[row][col] {
		return false
	}

	// add path
	(*path)[row][col] = true

	// continue dfs
	// down
	if existBT(board, row+1, col, word, wordIDX+1, path) {
		return true
	}
	// up
	if existBT(board, row-1, col, word, wordIDX+1, path) {
		return true
	}
	// right
	if existBT(board, row, col+1, word, wordIDX+1, path) {
		return true
	}
	// left
	if existBT(board, row, col-1, word, wordIDX+1, path) {
		return true
	}

	// remove from path
	(*path)[row][col] = false

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	//word := "ABCCED"
//	//word := "SEE"
//	word := "ABCB"
//	board1 := []byte("ABCE")
//	board2 := []byte("SFCS")
//	board3 := []byte("ADEE")
//	board := make([][]byte, 0)
//	board = append(board, board1)
//	board = append(board, board2)
//	board = append(board, board3)
//	fmt.Println(exist(board, word))
//}

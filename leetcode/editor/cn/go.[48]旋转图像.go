/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022-02-02 19:13:58
 * @version     v1.0
 * @filename    go.[48]旋转图像
 * @description
 ***************************************************************************/
package main

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]
//
//
// 示例 2：
//
//
//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
//
//
//
// 提示：
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
//
//
//
// Related Topics 数组 数学 矩阵 👍 1149 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func rotate(matrix *[][]int) {
//	res := [][]int{}
//	tmp := []int{}
//	for col := 0; col <= len(*matrix)-1; col++ {
//		tmp = []int{}
//		for row := len(*matrix) - 1; row >= 0; row-- {
//			tmp = append(tmp, (*matrix)[row][col])
//		}
//		res = append(res, tmp)
//	}
//	copy(*matrix, res)
//}

func rotate(matrix [][]int) {
	for inf, sup := -1, len(matrix); sup-inf > 2; {
		inf++
		sup--
		for i := 0; i < sup-inf; i++ {
			inf_in, sup_in := inf+i, sup-i

			//rotate
			matrix[inf][inf_in], matrix[inf_in][sup],
				matrix[sup_in][inf], matrix[sup][sup_in] =

				matrix[sup_in][inf], matrix[inf][inf_in],
				matrix[sup][sup_in], matrix[inf_in][sup]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
//	rotate(matrix)
//	fmt.Println(matrix)
//}

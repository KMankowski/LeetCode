package main

import "fmt"

func main() {
	fmt.Println(minFallingPathSum(nil))
}

func minFallingPathSum(matrix [][]int) int {

	if matrix == nil {
		return 0
	}
	for row := range matrix {
		if matrix[row] == nil {
			return 0
		}
	}

	dpMatrix := make([][]int, len(matrix))
	for row := range dpMatrix {
		dpMatrix[row] = make([]int, len(matrix[row]))
	}
	for col := range dpMatrix[0] {
		dpMatrix[0][col] = matrix[0][col]
	}

	for row := range matrix {
		if row == 0 {
			continue
		}
		for col := range matrix[row] {
			dpMatrix[row][col] = matrix[row][col] + getMinParent(dpMatrix, row, col)
		}
	}

	return getMin(dpMatrix[len(dpMatrix)-1])
}

func getMinParent(dpMatrix [][]int, row, col int) int {
	minParent := dpMatrix[row-1][col]

	if col-1 >= 0 && dpMatrix[row-1][col-1] < minParent {
		minParent = dpMatrix[row-1][col-1]
	}

	if col+1 <= len(dpMatrix[row])-1 && dpMatrix[row-1][col+1] < minParent {
		minParent = dpMatrix[row-1][col+1]
	}

	return minParent
}

func getMin(row []int) int {
	min := row[0]
	for i := range row {
		if row[i] < min {
			min = row[i]
		}
	}
	return min
}

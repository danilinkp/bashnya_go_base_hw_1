package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Scan(&a)
			row[j] = a
		}
		matrix[i] = row
	}
	isSimmetric := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				isSimmetric = false
			}
		}
	}
	if isSimmetric {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}

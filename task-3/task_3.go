package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	elem_end := arr[n-1]
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = elem_end
	fmt.Println(arr)
}

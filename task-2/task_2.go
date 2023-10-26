package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%d\n", &c)
	if a+b < c || a+c < b || b+c < a {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}

}

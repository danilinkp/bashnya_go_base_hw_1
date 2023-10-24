package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
	fmt.Println(c)
}

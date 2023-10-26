package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	numbers := strings.Fields(s)

	cAll := 0
	cContains := 0
	for i := 0; i < len(numbers); i++ {
		cAll++
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				cContains++
				break
			}
		}
	}
	fmt.Print(cAll - cContains)
}

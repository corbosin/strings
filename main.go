package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("nice try"))
	fmt.Println(count("nice try"))
}

func reverse(s string) string {
	slice := strings.Split(s, "")

	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return strings.Join(slice, "")
}

func count(s string) int {
	slice := strings.Split(s, "")
	n := 0
	for i := 0; i <= len(slice) - 1; i++ {
		n++
	}
	return n
}
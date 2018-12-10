package main

import (
	"fmt"
)

func max(e int, g int, k *int) {
	if e > g {
		*k = e
	} else {
		*k = g
	}
}
func main() {
	i := 10
	j := 14
	var c int
	max(i, j, &c)
	fmt.Println(c)
}

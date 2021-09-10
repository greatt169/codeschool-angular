package main

import (
	"fmt"
	"math/rand"
	"time"
)

func uniqRandn(n int) []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, n)
	dict := make(map[int]int)
	for i := 0; i < n; i++ {
		x := rand.Int()
		if _, ok := dict[x]; ok == false {
			result[i] = x
			dict[x]++
		} else {
			continue
		}
	}
	return result
}

func main() {
	fmt.Println(uniqRandn(10))
}

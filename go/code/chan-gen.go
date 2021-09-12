package main

import (
	"log"
	"math/rand"
	"time"
)

func genOneRandom(generated chan int, dict map[int]int) {

}

func uniqRandn(n int) []int {
	result := make([]int, n)
	dict := make(map[int]int)
	i := 0
	for i != n {
		rand.Seed(time.Now().UnixNano())
		x := rand.Int()
		if _, ok := dict[x]; !ok {
			result[i] = x
			dict[x]++
			i++
		} else {
			continue
		}
	}
	return result
}

func main() {
	start := time.Now()
	uniqRandn(10)
	elapsed := time.Since(start)
	log.Printf("uniqRandn took %s", elapsed)
}

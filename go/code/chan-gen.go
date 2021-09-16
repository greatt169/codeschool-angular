package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var mu = &sync.Mutex{}

func getRandEl(dict map[int]bool, chUniq chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	var x int
	mu.Lock()
	for {
		x = rand.Int()
		if !dict[x] {
			dict[x] = true
			break
		}
	}
	mu.Unlock()
	chUniq <- x
}

func uniqRandn(n int) []int {
	result := make([]int, n)
	dict := make(map[int]bool)
	chUniq := make(chan int, n)
	wg := sync.WaitGroup{}
	for i := 0; i != n; i++ {
		wg.Add(1)
		go getRandEl(dict, chUniq, &wg)
	}
	wg.Wait()
	close(chUniq)
	index := 0
	for v := range chUniq {
		result[index] = v
		index++
	}
	return result

}

func main() { // 15.655011ms
	start := time.Now()
	uniqRandn(1000)
	elapsed := time.Since(start)
	log.Printf("uniqRandn took %s", elapsed)
}

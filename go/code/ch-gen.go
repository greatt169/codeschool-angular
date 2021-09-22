package main

import (
	"fmt"
	"log"
	"time"
)

func fib(n int, c chan int) {
	for i, j, counter := 0, 1, 0; counter <= n; i, j, counter = i+j, i, counter+1 {
		c <- i
	}
	close(c)
}

func main() {

	start := time.Now()

	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// ..
	c := make(chan int, 10)
	go fib(10, c)
	for fn := range c {
		fmt.Println("Current fibonacci number is", fn)
	}
	elapsed := time.Since(start)
	log.Printf("uniqRandn took %s", elapsed)
}

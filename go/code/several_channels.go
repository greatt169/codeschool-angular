package main

import (
	"fmt"
	"sync"
)

func test(i int, imported chan int, errors chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if i < 6 {
		imported <- i
	} else {
		errors <- i
	}
}

func main() {
	total := 10
	imported := make(chan int, total)
	errors := make(chan int, total)
	wg := sync.WaitGroup{}
	for i := 0; i < total; i++ {
		wg.Add(1)
		go test(i, imported, errors, &wg)
	}
	wg.Wait()
	close(imported)
	close(errors)
	fmt.Println("good elements")
	for v1 := range imported {
		fmt.Println(v1)
	}
	fmt.Println("bad elements")
	for v2 := range errors {
		fmt.Println(v2)
	}
}

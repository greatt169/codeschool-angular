package main

import (
	"fmt"
	"sync"
)

func joinChannels(out chan int, chs ...chan int) {

	//go func() {
	wg := sync.WaitGroup{}
	wg.Add(len(chs))
	for _, ch := range chs {
		fmt.Println("go func(wg *sync.WaitGroup, ch chan int)")
		go func(wg *sync.WaitGroup, ch chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(&wg, ch)
	}
	wg.Wait()
	close(out)
	//}()
}

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	out := make(chan int)
	go joinChannels(out, a, b, c)
	for num := range out {
		fmt.Println(num)
	}
}

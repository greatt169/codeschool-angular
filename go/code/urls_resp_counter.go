package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkResponse(url string, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		c <- -1
		return
	}
	c <- resp.StatusCode
}

func main() {

	var with200 int
	var withNot200 int
	var withErr int

	urls := []string{"http://ozon.ru", "http://google.ru", "http://fb.com", "tt.tt", "https://www.asna.ru/111/"}
	var wg sync.WaitGroup
	c := make(chan int, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go checkResponse(url, &wg, c)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	for {
		v, ok := <-c
		if ok == false {
			fmt.Prin
			tln(with200)
			fmt.Println(withNot200)
			fmt.Println(withErr)
			return
		}
		switch v {
		case http.StatusOK:
			with200++
		case -1:
			withErr++
		default:
			withNot200++
		}

	}

}

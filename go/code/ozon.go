package main

import (
	"fmt"
)

func a() {
	x := []int{}
	x = append(x, 0)  // x = [0]
	x = append(x, 1)  // x = [0 1]
	x = append(x, 2)  // x = [0 1 2] // ключевой момент
	y := append(x, 3) // x = [0 1 2], y = [0 1 2 3]
	z := append(x, 4) // x = [0 1 2], y = [0 1 2 3], z = [0 1 2 4]
	fmt.Println(y, z)
}
func main() {
	//a()
	/**
	Что выведет программа?
	Ответ
	[0 1 2  4] [0 1 2  4]
	*/

	/*
		Что выведет программа?
	*/

	s := fmt.Sprintf("%d", -1999)
	fmt.Println(s)

	/*for i, _ := range s {
		fmt.Printf(fmt.Sprintf("%x ", s[i]))
	}
	fmt.Println(len([]rune(s)))*/
	//sn := "к" + string([]rune(s)[1:len(s) - 1])
	//s[0] = "R"
	//fmt.Println(s)
	//fmt.Println(sn)

	/*
		Требуется реализовать функцию uniqRandn, которая генерирует слайс длины n, состоящий из уникальных рандомных чисел.

		import (
		    "fmt"
		    "math/rand"
		)

		func main() {
		    fmt.Println(uniqRandn(10))
		}

		func uniqRandn(n int) []int {
		    sl := []int{}
		    dict := make(map[int]int)
		    for (i := 0; i < len(n); i++) {
		        if (len(dict[rand_d]) == 0) {
		            sl = append(sl, rand_n)
		            dict[rand_n] ++
		        }

		    }
		    //...
		}





	*/
}

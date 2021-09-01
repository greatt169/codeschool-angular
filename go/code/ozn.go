package main

import "fmt"

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}
func main() {
	a()
	/**
	Что выведет программа?

	func a() {
	    x := []int{}
	    x = append(x, 0)
	    x = append(x, 1)
	    x = append(x, 2)
	    y := append(x, 3)
	    z := append(x, 4)
	    fmt.Println(y, z)
	}

	func main() {
	    a()
	}


	x = 0 1 2 3
	y = 0 1 2 3
	z = 0 1 2 3 4


	Что выведет программа?

	func main() {
	    s := "тест"
	    fmt.Println(s[0])
	    sn := s[1, len(s)]
	    sn = "R" + sn
	    //s[0] = "R"
	    fmt.Println(sn)
	}


	t




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

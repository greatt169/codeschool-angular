package main

import "fmt"

/*
* На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение
 */
func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var res []int
	for _, v := range a {
		counter[v]++
	}
	for _, v := range b {
		if count, ok := counter[v]; ok && count > 0 {
			counter[v] -= 1
			res = append(res, v)
		}
	}
	return res
}

func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}

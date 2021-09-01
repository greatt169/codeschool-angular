package main

import "fmt"

const (
	zero = iota
	_
	two
)

func main() {
	fmt.Println(zero)
	//fmt.Println(one)
	fmt.Println(two)
}

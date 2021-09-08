package main

import "fmt"

func main() {
	v := 5
	p := &v
	println(*p)
	changePointer(p)
	println(*p)
}

func changePointer(p *int) {
	v := 3
	*p = v

	x := 1.5
	square(&x)
	fmt.Println(x)
}

func square(x *float64) {
	*x = *x * *x
}

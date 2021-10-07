package main

import "fmt"

type Test struct {
	Field int
}

func (t Test) ChangeF() {
	t.Field = 1
}

func main() {
	b := Test{}
	t := &b
	t.ChangeF()
	fmt.Println(t)
}

package main

import "fmt"

func main() {
	i := 1
	defer func(i *int) {
		fmt.Printf("%v ", *i)
	}(&i)
	i = 2
	defer func(i *int) {
		fmt.Printf("%v ", *i)
	}(&i)
}

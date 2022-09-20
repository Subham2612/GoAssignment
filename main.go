package main

import "fmt"

func main() {
	a, b := 10, -1

	if b > 0 {
		for b > 0 {
			a++
			b--
		}
	}

	if b < 0 {
		for b < 0 {
			a--
			b++
		}
	}

	fmt.Println(a)
}

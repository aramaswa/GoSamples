package main

import "fmt"

// function closure..  fibonacci is a function clusure
func fibonacci() func() int {
	a, b  := 0, 1
	return func () int {
		a = b + a
		b = a - b
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


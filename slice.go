package main

import (
	"fmt"
)

func main () {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1, 2, 3, 4, 5, 6, 7)
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	var fib = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i, v := range fib {
		fmt.Printf("[i,v] = [%d, %d]\n", i, v)
	}

	for _, v := range fib {
		fmt.Printf("[v] = [%d]\n", v)
	}

	for i := range fib {
		fmt.Println(i)
	}

	exp := make ([]int, 10)

	for i := range exp {
		exp[i] = 1 << uint8(i)
	}

	for _, v := range exp {
		fmt.Println( v)
	}
}

func printSlice(s string, slice []int) {
	fmt.Printf( "%s  len = %d  cap = %d   %v  \n ", s, len(slice), cap(slice), slice)
}


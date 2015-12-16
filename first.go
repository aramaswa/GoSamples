package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"math/cmplx"
)

func main() {
	defer fmt.Printf("\n Hello World... yayyyyy\n ")

	fmt.Println(time.Now())

	rand.Seed(20)
	fmt.Println("\n Some random number:", rand.Intn(67))

	fmt.Println("\n complex number:", cmplx.NaN())

	fmt.Println("\n 39989 + 7676 + 56= ", add(39989, 7676, 56))

	a,b := swap("hi", "anitha")

	fmt.Println( "\n Swap===== hi anitha ", a, b)

	fmt.Println(combine("hell", "life"))

	var yum rune

	fmt.Print(yum-1)

	fmt.Println("\n")
	fmt.Print(math.Sqrt(float64(68776)))
}


func add(x ,  y, z int) int {
	return x + y +z
}

func swap (s1, s2 string) (string, string) {
	return s2, s1
}

func combine (s1, s2 string) (s string, n int) {
	s = s1 + s2

	n = len(s)

	for i:=0 ; i <10 ; i++ {
		defer fmt.Println(i)
	}

/*	for  {

	}*/

	return // naked return.. you are returning without the names os the return variables. They are assumed from the value types
}
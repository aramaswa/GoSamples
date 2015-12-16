package main

import (
	"fmt"
)

type Point struct {
	X int
 	Y int
}

func main() {
	fmt.Println(Point{5,6})

	p := Point{8,5}
	addr := &p
	p.X = 87
	fmt.Println("address = ", addr)

	(*addr).Y = 98

	fmt.Println("address = ", addr)
	fmt.Println("pooint = ", p)
}

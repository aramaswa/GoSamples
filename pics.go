package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var p = make([][]uint8, dy)

	for i := range p {
		p[i] = make([]uint8, dx)
		for j := range p[i] {
			y := i ^ j
			p[i][j] = uint8(y)
		}

	}
	return p
}

func Foo(i, j int) int {
	fmt.Print("I'm hereeeeeeeeeeee")
	return i + j
}

func main() {
	pic.Show(Pic)

}

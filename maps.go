package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Vertex struct {
	Lat, Lon float64
}

var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)

	m["anitha"] = Vertex{
		45.34534,  56.4435,
	}

	fmt.Println(m["anitha"])

	var mm = map[string]Vertex{
		"abc" : Vertex{45.3, 56.7},
		"def" : Vertex{22.2, 9.3453456},
	}

	fmt.Println(mm)

	wc.Test(WordCount)

}

func WordCount(s string) map[string]int {
	var words = make (map[string]int)
	var t  = strings.Fields(s)
	for i :=0; i<len(t); i++{
		words[t[i]]++;
    }
	return words
}
package main
import (
	"fmt"
	"math"
	"os"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Hypo () float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyInt int32

func (v MyInt) double() int32 {
	return int32(v * 2)
}

func main() {
	p := &Vertex{3, 4}
	fmt.Println(p.Hypo())

	fmt.Println(MyInt(32).double())

	fmt.Fprintf(os.Stdout, "Heloo thereeeeeeeeeeeee")
}

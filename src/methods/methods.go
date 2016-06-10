package main

import (
	"geometry"
	"fmt"
)

func main() {
	p := geometry.Point{1,2}
	q := geometry.Point{3,4}
	fmt.Println(p.Distance(q))
	pointp := &p
	pointp.ScaleBy(50)
	fmt.Printf("%f, %f\n", p.X, p.Y)
}

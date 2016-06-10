package main

import (
	"geometry"
	"bitset"
	"fmt"
)

func main() {
	p := geometry.Point{1,2}
	q := geometry.Point{3,4}
	fmt.Println(p.Distance(q))
	pointp := &p
	pointp.ScaleBy(50)
	fmt.Printf("%f, %f\n", p.X, p.Y)

	var x, y bitset.BitSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
	x.Remove(9)
	fmt.Println(x.Len())
	fmt.Println(x.String())
  x.Clear()
	fmt.Println(x.String())
	fmt.Println(x.Len())
	x.Add(9)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	cpy := x.Copy()
	cpy.Add(100)
	fmt.Println(x.String())
	fmt.Println(cpy.String())
}

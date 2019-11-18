/* Hands-on exercise #5
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle */

package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
}
type CIRCLE struct {
	radius float64
}
func (c CIRCLE) area() float64 {
	return math.Pi * 2
}
func (s SQUARE) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func fArea (s shape) {
	fmt.Println(s.area())
}

func main() {

	s := SQUARE{5.5}
	c := CIRCLE{3.2}
	fArea(s)
	fArea(c)
}

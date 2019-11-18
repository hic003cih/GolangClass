/*
Method sets revisited
“The method set of a type determines the INTERFACES that the type implements.....” ~
golang spec
Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T
code: https://play.golang.org/p/KK8gjsAWBZ

*/

package main

import (
"fmt"
"math"
)

type circle struct {
radius float64
}

type shape interface {
area() float64
}

func (c *circle) area() float64 {
return math.Pi * c.radius * c.radius
}

func info(s shape) {
fmt.Println("area", s.area())
}

func main() {
c := circle{5}
// info(c)
fmt.Println(c.area())
}

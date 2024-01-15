package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

func (c circle) area() float64{
	return math.Pi * c.radius**2 
}

circle.radius = 64.23
fmt.Println(area())

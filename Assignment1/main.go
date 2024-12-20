package main

import (
	"Assignment1/Shapes"
	"fmt"
)

func main() {
	fmt.Println("----------------------------SHAPES-------------------------------")

	shapes := []Shapes.Shape{
		Shapes.Rectangle{Length: 5, Width: 3},
		Shapes.Circle{Radius: 4},
		Shapes.Square{Length: 6},
		Shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d:\n", i+1)
		Shapes.PrintShapeDetails(shape)
		fmt.Println()
	}
}

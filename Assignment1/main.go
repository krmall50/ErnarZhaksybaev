package main

import (
	"Assignment1/Employee"
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
	fmt.Println("----------------------------Employees-------------------------------")

	company := Employee.Company{Employees: make(map[string]Employee.Employee)}

	fte := Employee.FullTimeEmployee{ID: 1, Name: "Jessie", Salary: 59000}
	pte := Employee.PartTimeEmployee{ID: 2, Name: "Janet", HourlyRate: 3000, HoursWorked: 25.5}

	company.AddEmployee(fte)
	company.AddEmployee(pte)

	company.ListEmployees()
}

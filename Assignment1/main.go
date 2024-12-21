package main

import (
	//"Assignment1/Bank"
	"Assignment1/Employee"
	"Assignment1/Library"
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
	/*
		fmt.Println("----------------------------Employees-------------------------------")

		account := &Bank.BankAccount{
			AccountNumber: "123456789",
			HolderName:    "Jacky",
			Balance:       0.0,
		}

		for {
			fmt.Println("\nSelect an option:")
			fmt.Println("1. Deposit")
			fmt.Println("2. Withdraw")
			fmt.Println("3. Get Balance")
			fmt.Println("4. Exit")

			var choice int
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				fmt.Print("Enter amount to deposit: ")
				var amount float64
				fmt.Scanln(&amount)
				account.Deposit(amount)

			case 2:
				fmt.Print("Enter amount to withdraw: ")
				var amount float64
				fmt.Scanln(&amount)
				account.Withdraw(amount)

			case 3:
				account.GetBalance()

			case 4:
				fmt.Println("Exiting...")
				return

			default:
				fmt.Println("Invalid choice. Try again.")
			}
		}
	*/
	fmt.Println("----------------------------LIBRARY-------------------------------")

	lib := Library.Library{Books: make(map[string]Library.Book)}

	lib.AddBook(Library.Book{ID: "1", Title: "book1", Author: "ernar", IsBorrowed: false})
	lib.AddBook(Library.Book{ID: "2", Title: "book2", Author: "author", IsBorrowed: false})
	lib.AddBook(Library.Book{ID: "3", Title: "book4", Author: "author2", IsBorrowed: false})
	lib.AddBook(Library.Book{ID: "4", Title: "book3", Author: "author8", IsBorrowed: false})

	for {
		var input string
		fmt.Println("Choose an option: (Add, Borrow, Return, List, Exit)")
		fmt.Scanln(&input)

		switch input {
		case "Add":
			var id, title, author string
			fmt.Println("Enter Book id:")
			fmt.Scanln(&id)
			fmt.Println("Enter Book Title:")
			fmt.Scanln(&title)
			fmt.Println("Enter Book Author:")
			fmt.Scanln(&author)
			book := Library.Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			lib.AddBook(book)

		case "Borrow":
			var id string
			fmt.Println("Enter Book ID to borrow:")
			fmt.Scanln(&id)
			lib.BorrowBook(id)

		case "Return":
			var id string
			fmt.Println("Enter Book ID to return:")
			fmt.Scanln(&id)
			lib.ReturnBook(id)

		case "List":
			lib.ListBooks()

		case "Exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option")
		}
	}

}

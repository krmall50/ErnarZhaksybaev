package main

import (
	"fmt"
)

func main() {
	lib := Library{Books: make(map[string]Book)}

	lib.addBook(Book{ID: "1", Title: "book1", Author: "ernar", IsBorrowed: false})
	lib.addBook(Book{ID: "2", Title: "book2", Author: "author", IsBorrowed: false})
	lib.addBook(Book{ID: "3", Title: "book4", Author: "author2", IsBorrowed: false})
	lib.addBook(Book{ID: "4", Title: "book3", Author: "author8", IsBorrowed: false})

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
			book := Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			lib.addBook(book)

		case "Borrow":
			var id string
			fmt.Println("Enter Book ID to borrow:")
			fmt.Scanln(&id)
			lib.borrowBook(id)

		case "Return":
			var id string
			fmt.Println("Enter Book ID to return:")
			fmt.Scanln(&id)
			lib.returnBook(id)

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

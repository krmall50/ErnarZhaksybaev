package main

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

func silly() {
	fmt.Println("helo")
}

type Library struct {
	Books map[string]Book
}

func (l *Library) addBook(book Book) {
	l.Books[book.ID] = book
	fmt.Println("Book added: " + book.Title)
}
func (l *Library) borrowBook(id string) {
	book, exists := l.Books[id]
	if exists && !book.IsBorrowed {
		book.IsBorrowed = true
		l.Books[id] = book
		fmt.Println("Book borrowed: " + book.Title)
	} else if exists {
		fmt.Println("book is borrowed: " + book.Title)
	} else {
		fmt.Println("Book is not found.")
	}
}
func (l *Library) returnBook(id string) {
	book, exists := l.Books[id]
	if exists && book.IsBorrowed {
		book.IsBorrowed = false
		l.Books[id] = book
		fmt.Println("Book returned:", book.Title)
	} else if exists {
		fmt.Println("Book is not borrowed:", book.Title)
	} else {
		fmt.Println("Book isn't found.")
	}
}
func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

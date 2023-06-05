// Book Struct: Design a struct for a Book with fields like Title, Author, Pages, and Publication Year. Write a function that
// takes a slice of Books and calculates the total number of pages.

package main

import (
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Pages           uint
	PublicationYear uint
}

func TotalPageCounter(bookList []Book) uint {
	pageCounter := 0
	for _, book := range bookList {
		pageCounter += int(book.Pages)
	}

	return uint(pageCounter)
}

func main() {
	book1 := Book{
		Title:           "Angels & Demon's",
		Author:          "Dan Brown",
		Pages:           768,
		PublicationYear: 2000,
	}

	book2 := Book{
		Title:           "The Da Vinci Code",
		Author:          "Dan Brown",
		Pages:           689,
		PublicationYear: 2003,
	}

	book3 := Book{
		Title:           "The Lost Symbol",
		Author:          "Dan Brown",
		Pages:           528,
		PublicationYear: 2009,
	}

	book4 := Book{
		Title:           "Inferno",
		Author:          "Dan Brown",
		Pages:           642,
		PublicationYear: 2013,
	}

	book5 := Book{
		Title:           "Origin",
		Author:          "Dan Brown",
		Pages:           461,
		PublicationYear: 2017,
	}

	bookList := []Book{book1, book2, book3, book4, book5}

	totalNumberOfPages := TotalPageCounter(bookList)

	fmt.Printf("Total number of pages in all Dan Bown's books about Robert Langdon is: %v pages\n", totalNumberOfPages)

}

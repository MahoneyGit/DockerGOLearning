package book

import (
	"fmt"
)

type Book struct {
	BookID          int
	name            string
	numOfPages      int
	bookDescription string
	publisherId     int
}

func CreateBook(bookName string, numOfPages int, bookDescription string, publisherId int) (Book, error) {
	b := Book{
		name:            bookName,
		numOfPages:      numOfPages,
		bookDescription: bookDescription,
		publisherId:     publisherId,
	}
	fmt.Println("Book created, details are as follows, if incorrect you are fucked")
	b.PrintDetails()
	return b, nil
	// I could have created the book via new(Book) but this creates B as an object pointer. Since I am returning a value instead
	// I was confused how I can still access PrintDetails but Go allows method promotion
}

// Todo figure out how convert book from db to struct
func GetBookByID(bookID int) (*Book, error) {
	fmt.Println("Finding book..... \nActing Busy.... \nNot yet implemented.... ", bookID)
	return &Book{}, nil
}

func DeleteBookByID(bookID int) (bool, error) {
	fmt.Println("Deleting book..... ", bookID)
	return true, nil
}

func UpdateBook(bookID int) (bool, error) {
	fmt.Println("Creating book in Database..... ", bookID)
	fmt.Println("Book Created..... ")
	return true, nil
}

func (book *Book) PrintDetails() {
	fmt.Println("book name is: ", book.name)
	fmt.Println("book page num is: ", book.numOfPages)
	fmt.Println("book description is: ", book.bookDescription)
	fmt.Println("book publisherId is: ", book.publisherId)
}

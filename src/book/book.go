package book

import "fmt"

type Book struct {
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

func (book *Book) PrintDetails() {
	fmt.Println("book name is: ", book.name)
	fmt.Println("book page num is: ", book.numOfPages)
	fmt.Println("book description is: ", book.bookDescription)
	fmt.Println("book publisherId is: ", book.publisherId)
}

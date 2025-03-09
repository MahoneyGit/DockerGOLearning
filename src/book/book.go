package book

import (
	"fmt"
)

type Book struct {
	BookID          int    `json:"book_id,omitempty" validate:"omitempty"`
	Name            string `json:"bookName,omitempty" validate:"required,min=3"`
	NumOfPages      int    `json:"num_of_pages,omitempty" validate:"required,min=1"`
	BookDescription string `json:"book_description,omitempty" validate:"required"`
	PublisherId     int    `json:"publisher_id,omitempty" validate:"omitempty"`
}

func CreateBook(bookName string, numOfPages int, bookDescription string, publisherId int) (Book, error) {
	b := Book{
		Name:            bookName,
		NumOfPages:      numOfPages,
		BookDescription: bookDescription,
		PublisherId:     publisherId,
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

func UpdateBook(book Book) (bool, error) {
	fmt.Println("Creating book in Database..... ", book.BookID)
	fmt.Println("Book Created..... ")
	return true, nil
}

func (book *Book) PrintDetails() {
	fmt.Println("book name: ", book.Name)
	fmt.Println("book page num: ", book.NumOfPages)
	fmt.Println("book description: ", book.BookDescription)
	fmt.Println("book publisherId: ", book.PublisherId)
}

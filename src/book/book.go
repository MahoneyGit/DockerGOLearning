package book

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/db"
)

type Book struct {
	BookID          int    `json:"book_id,omitempty" validate:"omitempty"`
	Name            string `json:"book_name,omitempty" validate:"required,min=3"`
	NumOfPages      int    `json:"num_of_pages,omitempty" validate:"required,min=1"`
	BookDescription string `json:"book_description,omitempty" validate:"required"`
	PublisherId     int    `json:"publisher_id,omitempty" validate:"omitempty"`
}

func CreateBook(bookName string, numOfPages int, bookDescription string, publisherId int) (*Book, error) {
	b := Book{
		Name:            bookName,
		NumOfPages:      numOfPages,
		BookDescription: bookDescription,
		PublisherId:     publisherId,
	}
	fmt.Println("Book created, details are as follows, if incorrect you are fucked")
	b.PrintDetails()
	return &b, nil
	// I could have created the book via new(Book) but this creates B as an object pointer. Since I am returning a value instead
	// I was confused how I can still access PrintDetails but Go allows method promotion
}

func CreateBookFromSqlRow(rowResult *sql.Rows) (*Book, error) {
	b := &Book{}
	if rowResult.Next() {
		err := rowResult.Scan(&b.BookID, &b.Name, &b.NumOfPages, &b.BookDescription, &b.PublisherId)
		if err != nil {
			fmt.Printf("\n\nSomething has gone seriously wrong!%v\n\n", err)
			return nil, err
		}
	} else {
		return b, fmt.Errorf("no rows found")
	}
	fmt.Println("Book created, details are as follows, if incorrect you are fucked")
	b.PrintDetails()
	return b, nil
	// I could have created the book via new(Book) but this creates B as an object pointer. Since I am returning a value instead
	// I was confused how I can still access PrintDetails but Go allows method promotion
}

func (book *Book) SaveBook() (err error) {
	book.PublisherId = getPublisherId()
	saveBookQuery := fmt.Sprintf(`INSERT INTO public.book (book_name, num_of_pages, book_description, publisher_id)
	VALUES ('%s', %d, '%s', %d);`, book.Name, book.NumOfPages, book.BookDescription, book.PublisherId)
	fmt.Println(saveBookQuery)
	result, er := db.RunQuery(saveBookQuery)
	if er != nil {
		return er
	}
	fmt.Println(result)
	return nil
}

func GetBookByID(bookID int) (string, error) {
	retrieveBookQuery := fmt.Sprintf(`SELECT * FROM public.book WHERE book_id = %d;`, bookID)
	fmt.Println(retrieveBookQuery)
	// var retrievedBook *Book
	fmt.Println("Finding book..... \nActing Busy.... ", bookID)
	bookAsString, err := db.RunQuery(retrieveBookQuery)
	if err != nil {
		return "", err
	}

	retrievedBook, err := CreateBookFromSqlRow(bookAsString)
	if err != nil {
		return "", err
	}
	returnValue, err := json.Marshal(retrievedBook)
	fmt.Println(returnValue)
	return (string(returnValue)), nil
}

func DeleteBookByID(bookID int) (bool, error) {
	deleteBookQuery := fmt.Sprintf(`DELETE FROM public.book WHERE
	book_id = %d;`, bookID)
	fmt.Println(deleteBookQuery)
	fmt.Println("Deleting book..... ", bookID)
	result, err := db.RunQuery(deleteBookQuery)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return true, nil
}

func (book *Book) UpdateBook() (bool, error) {
	fmt.Println("Updating book in Database..... ", book.BookID)
	book.PublisherId = getPublisherId()
	saveBookQuery := fmt.Sprintf(`UPDATE public.book 
	SET
    	book_name = '%s',
    	num_of_pages = %d,
    	book_description = '%s',
    	publisher_id = %d
	WHERE
    	book_id = %d;`,
		book.Name, book.NumOfPages, book.BookDescription, book.PublisherId, book.BookID)
	fmt.Println(saveBookQuery)
	result, err := db.RunQuery(saveBookQuery)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	fmt.Println("Book Updated..... ")
	return true, nil
}

func (book *Book) PrintDetails() {
	fmt.Println("book name: ", book.Name)
	fmt.Println("book page num: ", book.NumOfPages)
	fmt.Println("book description: ", book.BookDescription)
	fmt.Println("book publisherId: ", book.PublisherId)
}

func getPublisherId() int {
	fmt.Printf("Generating new Random Number\n")

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(101) // The upper bound is exclusive, so 101 gives us numbers from 0 to 100
}

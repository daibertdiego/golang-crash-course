package main

import (
	"errors"
	"fmt"
)

func main() {

}

func Add(num1, num2 float64) (float64, error) {
	result := num1 + num2
	if result > 15 {
		return 0, errors.New("result number too large for this function")
	}
	return result, nil
}

type LibraryService struct {
	DB LibraryDB
}

func NewLibraryService(db LibraryDB) *LibraryService {
	return &LibraryService{
		DB: db,
	}
}

func (service LibraryService) CheckOutBook(bookID string, userID string) error {
	book := service.DB.GetBook(bookID)
	if book.CheckOutByUserID != nil {
		return fmt.Errorf("book with id %s is currently checked out", bookID)
	}
	return service.DB.CheckOutBook(bookID, userID)
}

type LibraryDB interface {
	GetBook(bookID string) DBBook
	CheckOutBook(bookID string, userID string) error
}

type DBBook struct {
	ID               string
	CheckOutByUserID *string
}

package dao

import (
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书是%v", k, v)
	}
}

//func TestAddBooks(t *testing.T) {
//	book := &model.Book{
//		Title:   "sgyy",
//		Author:  "luo",
//		Price:   "59.99",
//		Sales:   "100",
//		Stock:   "10",
//		ImgPath: "/static/img/default.jpg",
//	}
//	AddBook(book)
//}

func TestDeleteBooks(t *testing.T) {
	DeleteBook("1")
}

func TestGetBookByID(t *testing.T) {
	book, _ := GetBookByID("8")
	fmt.Println("book = ", book)
}

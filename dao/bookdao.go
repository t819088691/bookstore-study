package dao

import (
	"bookstore-study/model"
	"bookstore-study/utils"
)

func GetBooks() ([]*model.Book, error) {
	sql := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil

}

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

func DeleteBook(bookID string) error {
	sql := "delete from books where id = ?"
	_, error := utils.Db.Exec(sql, bookID)
	if error != nil {
		return error
	}
	return nil
}


func GetBookByID(bookID string) (*model.Book, error) {
	sql := "sid,title,author,price,sales,stock,img_path from books where ID = ?"
	row := utils.Db.QueryRow(sql, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

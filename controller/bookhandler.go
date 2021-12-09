package controller

import (
	"bookstore-study/dao"
	"net/http"
	"text/template"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

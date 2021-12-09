package main

import (
	"bookstore-study/controller"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/main", IndexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	http.ListenAndServe(":8080", nil)
}

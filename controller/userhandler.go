package controller

import (
	"bookstore-study/dao"
	"net/http"
	"text/template"
)

//login
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

//regist
func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}
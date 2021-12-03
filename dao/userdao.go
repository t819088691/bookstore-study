package dao

import (
	"bookstore-study/model"
	"bookstore-study/utils"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {

	sql := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {

	sql := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, err
}

func SaveUser(username string, password string, email string) error {

	sql := "insert into users (username,password,email) value (?,?,?)"
	_, err := utils.Db.Exec(sql, username, password, email)
	if err != nil {
		return err
	}
	return nil

}

package dao

import (
	"fmt"
	"testing"
)

func TestCheckUserNameAndPassword(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("user =", user)
}

func TestCheckUserName(t *testing.T) {
	user, _ := CheckUserName("admin1")
	fmt.Println("user = ", user)
}

func TestSaveUser(t *testing.T) {
	SaveUser("admin1", "123456", "admin@qq.com")
}

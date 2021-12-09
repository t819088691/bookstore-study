package dao

import (
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%V本书是%V", k, v)
	}
}

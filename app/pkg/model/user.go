package model

import "github.com/mahiro72/go-mvc-server/pkg/database"

// import "github.com/jinzhu/gorm"

type User struct {
	Id   int
	Name string
}

func GetUser(id int) (*User,error) {
	var user User
	database.DB.First(user).Where("id = ?",id)
	return &user,nil
}

func CreateUser(name string) (*User,error) {
	user := &User{Name: name}
	database.DB.Create(user)
	return user,nil
}

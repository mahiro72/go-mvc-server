package model

import "github.com/mahiro72/go-mvc-server/pkg/database"

// import "github.com/jinzhu/gorm"

type User struct {
	Id   int
	Name string
}

func GetUser(id int) (*User,error) {
	var user User
	err := database.DB.Where("id = ?",id).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user,nil
}

func CreateUser(name string) (*User,error) {
	user := &User{Name: name}
	err := database.DB.Create(user).Error
	if err != nil {
		return nil,err
	}
	return user,nil
}

package model

import "github.com/mahiro72/go-mvc-server/pkg/persistence"

type User struct {
	Id   int
	Name string
}

func GetUser(id int) (*User, error) {
	var user User
	if err := persistence.DB.Where("id = ?", id).First(&user).Error;err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(name string) (*User, error) {
	user := &User{Name: name}
	if err := persistence.DB.Create(user).Error;err != nil {
		return nil, err
	}
	return user, nil
}

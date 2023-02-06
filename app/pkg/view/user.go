package view

import "github.com/mahiro72/go-mvc-server/pkg/model"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserToJSON(u *model.User) *User {
	return &User{
		ID:   u.ID,
		Name: u.Name,
	}
}

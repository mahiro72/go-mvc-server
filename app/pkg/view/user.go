package view

import "github.com/mahiro72/go-mvc-server/pkg/model"

type UserJSON struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UserToJSON(u *model.User) *UserJSON {
	return &UserJSON{
		Id:   u.Id,
		Name: u.Name,
	}
}

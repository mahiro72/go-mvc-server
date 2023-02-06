package view

import "github.com/mahiro72/go-mvc-server/pkg/model"

type UserJSON struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserToJSON(u *model.User) *UserJSON {
	return &UserJSON{
		ID:   u.ID,
		Name: u.Name,
	}
}

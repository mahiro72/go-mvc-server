package model

// import "github.com/jinzhu/gorm"

type User struct {
	Id   int
	Name string
}

func GetUser(id string) (*User,error) {
	// gorm
	return &User{Id: 1,Name: "hoge"},nil
}

func CreateUser(name string) (*User,error) {
	// gorm
	return &User{Id: 1,Name: name},nil
}



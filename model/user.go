package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int
	Name string
	Email string
	Password string
	Profile string
	ImageUrl string
}

func GetAllUsers() (datas []User) {
	result := Db.Find(&datas)
	fmt.Println(datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
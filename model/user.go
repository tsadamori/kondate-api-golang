package model

import (
	"github.com/tsadamori/kondate-api-golang/auth"
)

type User struct {
	Id uint
	Name string
	Email string
	Profile string
	ImageUrl string
}

type UserToRegister struct {
	Name string
	Email string
	Password string
	Profile string
	ImageUrl string
}

func GetAllUsers() (data []User) {
	result := Db.Find(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetUserById(id string) (data User) {
	result := Db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetUserPasswordByEmail(email string) (password string) {
	var data User
	result := Db.Model(&data).Select("password").Where("email = ?", email).First(&password)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func IsValidUser(email string, password string) bool {
	hash := GetUserPasswordByEmail(email)
	if auth.CompareHashAndPassword(hash, password) != nil {
		return false
	}
	return true
}
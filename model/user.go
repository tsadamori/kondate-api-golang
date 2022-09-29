package model


type User struct {
	Id uint
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
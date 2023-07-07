package model

type Category struct {
	Id uint
	Name string
}

func GetAllCategories(limit, offset int) (data []Category){
	result := Db.Limit(limit).Offset(offset).Find(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
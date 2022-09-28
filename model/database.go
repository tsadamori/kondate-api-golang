package model

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var user, password, db_host, db_port, db_name string
	var err error
	
	loadEnv(&user, &password, &db_host, &db_port, &db_name)

	var path string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, db_host, db_port, db_name)
	dialector := mysql.Open(path)
	fmt.Println(path)

	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!")
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second)
			count--
			fmt.Printf("retry... count%v\n", count)
			connect(dialector, count)
			return
		}
	}
	panic(err.Error())
}

func loadEnv(user *string, password *string, db_host *string, db_port *string, db_name *string) {
	godotenv.Load(".env")
	*user = os.Getenv("MYSQL_USER")
	*password = os.Getenv("MYSQL_PASSWORD")
	*db_port = os.Getenv("MYSQL_PORT")
	*db_host = os.Getenv("MYSQL_HOST")
	*db_name = os.Getenv("MYSQL_DATABASE")
}

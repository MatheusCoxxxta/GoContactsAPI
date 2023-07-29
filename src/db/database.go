package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	Models "api/src/models"
)

var CONNECTION *gorm.DB
var err error

const DNS = "root:password@tcp(127.0.0.1:3309)/contacts?charset=utf8mb4&parseTime=true&loc=Local"

type Contact Models.Contact
type User Models.User

func InitializeDatabase() {
	CONNECTION, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Not initialized")
	}
	CONNECTION.AutoMigrate(&User{}, &Contact{})
	fmt.Println("Migrated successfully...")
}

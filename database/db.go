package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDataBase() {
	dsn := "root:root@tcp(localhost:3306)/ApiGolang?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Panic("failed to connect database")
	}
}

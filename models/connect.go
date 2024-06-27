package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/todowleest_db"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Todo{})
	DB = database
}

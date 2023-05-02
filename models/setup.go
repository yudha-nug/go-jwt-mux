package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:skywalker@tcp(192.168.1.6:3306)/go-jwt-mux"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db

}
package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:P@ssw0rd123@tcp(localhost:3306)/local-go"))
	if err != nil {
		fmt.Println("Database connection failed")
	}

	db.AutoMigrate(&User{}, &Product{})

	DB = db

}

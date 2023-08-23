package services

import (
	"fmt"
	"main/configs"
	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	dbConfig := configs.MYSQL_USERNAME + ":" + configs.MYSQL_PASSWORD + "@tcp(" + configs.MYSQL_HOST + ":" + configs.MYSQL_PORT + ")/" + configs.MYSQL_NAME + "?parseTime=true&loc=Local"

	fmt.Println("Database Configuration: ", dbConfig)

	var err error
	Db, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: error=%v", err)
		return nil
	}

	return Db
}

func SyncDB() {
	Db.AutoMigrate(&models.User{})
}

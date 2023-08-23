package dao

import (
	"fmt"
	"main/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
  dbConfig := config.MYSQL_USERNAME + ":" + config.MYSQL_PASSWORD + "@tcp(" + config.MYSQL_HOST + ":" + config.MYSQL_PORT + ")/" + config.MYSQL_NAME + "?parseTime=true&loc=Local"
	
  fmt.Println("Database Configuration: ", dbConfig)

  var err error
  Db, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
  
	if err != nil {
		fmt.Printf("Error connecting to database: error=%v", err)
		return nil
	}
  
	return Db
}

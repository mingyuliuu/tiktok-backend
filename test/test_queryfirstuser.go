package test

import (
	"fmt"
	"main/dao"
)

var Q dao.Query

func QueryFirstUser() {
	db := dao.InitDB()
	Q = *dao.Use(db)
	user, err := Q.Users.First()
	if err != nil {
		fmt.Printf("Error querring the first user in database: error=%v", err)
		return
	}
	fmt.Println(user.UserID, user.UserName, user.UserPassword)
}

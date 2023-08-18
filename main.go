package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/dao"
)

func main() {
	// Initiate project dependencies
	dao.InitDB()

	r := gin.Default()
	initRouter(r)
	r.Run(":8080")

	fmt.Println("Application running on port 8080")
}

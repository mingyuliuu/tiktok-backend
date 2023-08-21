package main

import (
	"fmt"
	"main/dao"
	"main/middleware/gormgen"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initiate project dependencies

	// Init database
	dao.InitDB()

	// Automatically generate DAO from models by GEN
	gormgen.DaoGenerator()

	r := gin.Default()

	initRouter(r)

	r.Run(":8080")

	fmt.Println("Application running on port 8080")
}

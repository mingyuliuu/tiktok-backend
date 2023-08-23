package main

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
	"net/http"
)

func initRouter(r *gin.Engine) {
	r.GET("/main", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection - success!",
		})
	})

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
		userRouter.GET("/", controllers.GetUserInfo)
	}

	internalRouter := r.Group("/internal")
	{
		internalRouter.GET("/users", controllers.GetAllUsers)
	}

	/*
		  r.GET("/feed/", controller.Feed)

			r.POST("/user/register/", controller.Register)
			r.POST("/user/login/", controller.Login)
		  r.GET("/user/", controller.UserInfo)

			r.POST("/publish/action/", controller.PublishVideo)
			r.GET("/publish/list/", controller.GetPublishedVideos)
	*/
}

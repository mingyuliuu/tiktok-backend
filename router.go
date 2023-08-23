package main

import (
	"github.com/gin-gonic/gin"
	"main/controller"
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
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.GET("/", controller.GetUserInfo)
	}

	internalRouter := r.Group("/internal")
	{
		internalRouter.GET("/users", controller.GetAllUsers)
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

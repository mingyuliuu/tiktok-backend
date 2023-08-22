package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
  "main/controller"
)

func initRouter(r *gin.Engine) {
  r.GET("/main", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Connection - success!",
    })
  })

  r.POST("/user/register/", controller.Register)
  
  /*
  r.GET("/feed/", controller.Feed)

	r.POST("/user/register/", controller.Register)
	r.POST("/user/login/", controller.Login)
  r.GET("/user/", controller.UserInfo)
 
	r.POST("/publish/action/", controller.PublishVideo)
	r.GET("/publish/list/", controller.GetPublishedVideos)
  */
  
}

package controller

import (
	"main/dao"
  "main/models"
	"net/http"

  "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
  var body struct {
    Username string
    Password string
  }

  if c.Bind(&body) != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Failed to read body",
    })
    return
  }

  //Create 
  hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
  if err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to hash password",
    })
    return
  }

  // Create user
  user := models.User{Username: body.Username, Password: string(hash)}
  result := dao.Db.Create(&user)
  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "Failed to create user",
    })
  }

  // Respond
  c.JSON(http.StatusOK, gin.H{ "message": "User registered successfully"})
}

  //  POST request douyin/user/login/ 用户登录
  func Login(c *gin.Context){
    
  }

  //  GET request douyin/user/ 用户信息
  func UserInfo(c *gin.Context){
    
  }

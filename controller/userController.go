package controller

import (
	"main/dao"
	"main/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRegistrationResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token"`
}

func Register(c *gin.Context) {
	// Get info from body
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, UserRegistrationResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserRegistrationResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to hash password",
		})

		return
	}

	// Create user
	user := models.User{Username: body.Username, Password: string(hash)}
	result := dao.Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, UserRegistrationResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, UserRegistrationResponse{
		StatusCode: 0,
		// UserId: ,
		// Token: ,
	})
}

// POST request douyin/user/login/ 用户登录
func Login(c *gin.Context) {
	// Get info from body
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

	// Look up requested user
	var user models.User
	dao.Db.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare with the password hass in db
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//Sign and get encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)
	c.JSON(http.StatusOK, gin.H{})
}

// GET request douyin/user/ Get user info by ID
func UserInfo(c *gin.Context) {
	/*
	   var users []models.User
	   // Get all records
	   result := dao.Db.Find(&users)
	*/
	/*
	   user_id := c.Query("user_id")
	   id, _ := strconv.ParseInt(user_id, 10, 64)

	   if u, err := usi.GetUserById(id); err != nil {
	       c.JSON(http.StatusOK, )
	   } else {
	       c.JSON(http.StatusOK, UserResponse{
	           Response: Response{StatusCode: 0},
	           User:     u,
	       })
	   }
	*/
}

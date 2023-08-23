package controllers

import (
	"main/models"
	"main/services/database"
	"net/http"
	"os"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token"`
}

type UserInfoResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg,omitempty"`
	User       *models.User `json:"user"`
}

type GetUsersResponse struct {
	StatusCode int32         `json:"status_code"`
	Users      []models.User `json:"users"`
}

func Register(c *gin.Context) {
	var queryParams struct {
		Username string
		Password string
	}

	if c.Bind(&queryParams) != nil {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to read query parameters",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(queryParams.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to hash password",
		})

		return
	}

	// Create user
	newUser := models.User{Username: queryParams.Username, Password: string(hash)}
	result := services.Db.Create(&newUser)

	// Get user ID
	var findUser models.User
	services.Db.First(&findUser, "username = ?", queryParams.Username)
	userID := int64(findUser.ID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, UserAuthResponse{
		StatusCode: 0,
		UserId:     userID,
		Token:      "",
	})
}

func Login(c *gin.Context) {
	var queryParams struct {
		Username string
		Password string
	}

	if c.Bind(&queryParams) != nil {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to read query parameters",
		})

		return
	}

	// Look up requested user
	var user models.User
	services.Db.First(&user, "username = ?", queryParams.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Invalid email or password",
		})

		return
	}

	// Compare with the password hash in db
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(queryParams.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Invalid email or password",
		})

		return
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, UserAuthResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)

	c.JSON(http.StatusOK, UserAuthResponse{
		StatusCode: 0,
		UserId:     int64(user.ID),
		Token:      tokenString,
	})
}

func GetUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	var user models.User
	result := services.Db.First(&user, userID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to obtain user information",
			User:       nil,
		})

		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: 0,
		User:       &user,
	})
}

// For internal use: get the list of all users
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := services.Db.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, GetUsersResponse{
			StatusCode: 1,
		})

		return
	}

	c.JSON(http.StatusOK, GetUsersResponse{
		StatusCode: 0,
		Users:      users,
	})
}

package controller

import (
	"main/dao"
	"main/models"
	"net/http"
	"os"

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
type ListUsersResponse struct {
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
	result := dao.Db.Create(&newUser)

	// Get user ID
	var findUser models.User
	dao.Db.First(&findUser, "username = ?", queryParams.Username)
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
	dao.Db.First(&user, "username = ?", queryParams.Username)

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

// For internal use: get the list of all users
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := dao.Db.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, ListUsersResponse{
			StatusCode: 1,
		})

		return
	}

	c.JSON(http.StatusOK, ListUsersResponse{
		StatusCode: 0,
		Users:      users,
	})
}

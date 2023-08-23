package services

import (
	"log"
	"main/models"
	services "main/services/database"
)

type UserInfo struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
	WorkCount       int64  `json:"work_count,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
}

// Insert a new user into MySQL database and return its id
func CreateUser(user *models.User) (int64, error) {
	if err := services.Db.Create(&user).Error; err != nil {
		log.Println("Creating user error:", err.Error())
		return 0, err
	}

	var findUser models.User
	services.Db.First(&findUser, "username = ?", user.Username)

	return int64(findUser.ID), nil
}

// Get the list of all users from MySQL database
func GetUsersList() ([]models.User, error) {
	users := []models.User{}

	if err := services.Db.Find(&users).Error; err != nil {
		log.Println("Getting users list error:", err.Error())
		return users, err
	}

	return users, nil
}

// Get a user from MySQL database by its user_id
func GetUserByUsername(username string) (models.User, error) {
	user := models.User{}

	if err := services.Db.Where("username = ?", username).First(&user).Error; err != nil {
		log.Println("Getting user by username error:", err.Error())
		return user, err
	}

	return user, nil
}

// Get a user from MySQL database by its user_id
func GetUserById(id int64) (models.User, error) {
	user := models.User{}

	if err := services.Db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("Getting user by id error:", err.Error())
		return user, err
	}

	return user, nil
}

// Get the decorated UserInfo object by user_id
func GetUserInfoById(id int64) (*UserInfo, error) {
	rawUser, err := GetUserById(id)

	if err != nil {
		log.Println("Getting user info by id error:", err.Error())
		return nil, err
	}

	user := UserInfo{
		Id:              id,
		Name:            rawUser.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}

	return &user, nil
}

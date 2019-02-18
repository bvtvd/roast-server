package models

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	. "roast-server/database"
	"time"
)


type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`	
	Name    string  `json:"name"`
	Username string `json:username`
	Password string `json:password`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Likes []Cafe `gorm:"many2many:users_cafes_likes"`
}

// 登录获取
func GetUserFromCredentials(user User) (User, error) {	

	if err := Db.Where("username = ? and password = ?", user.Username, user.Password).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return user, err
	}

	return user, nil
}

// 根据id获取用户信息
func GetUserById(id int) User {
	var user User

	Db.First(&user, id)

	return user
}
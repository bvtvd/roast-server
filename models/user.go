package models

import (
	// "fmt"
	// "github.com/jinzhu/gorm"
	// . "roast-server/database"
	"time"
)


type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`	
	Name    string  `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Likes []Cafe `gorm:"many2many:users_cafes_likes"`
}
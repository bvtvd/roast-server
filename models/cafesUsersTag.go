package models

import (
	// "fmt"
	// "github.com/jinzhu/gorm"
	// . "roast-server/database"
	"time"
)


type CafesUsersTag struct {
	CafeId   uint  `gorm:"primary_key;type:int(11)" json:"cafe_id"`	
	UserId   uint  `gorm:"primary_key;type:int(11)" json:"user_id"`
	TagId    uint  `gorm:"primary_key;type:int(11)" json:"tag_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
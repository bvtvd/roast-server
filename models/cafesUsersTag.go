package models

import (
	// "fmt"
	// "github.com/jinzhu/gorm"
	. "roast-server/database"
	"time"
)


type CafesUsersTag struct {
	CafeId   uint  `gorm:"primary_key;type:int(11)" json:"cafe_id"`	
	UserId   uint  `gorm:"primary_key;type:int(11)" json:"user_id"`
	TagId    uint  `gorm:"primary_key;type:int(11)" json:"tag_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// 删除关联记录
func DeleteCafeTag(cafeId uint, tagId uint, userId uint) bool {
	var record CafesUsersTag = CafesUsersTag{CafeId: cafeId, UserId: userId, TagId: tagId}

	Db.Delete(&record)

	return true
}
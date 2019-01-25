package models

import (
	"github.com/jinzhu/gorm"
	. "roast-server/database"
	"fmt"
)

type Cafe struct {
	gorm.Model
	Name string	
	Address string
	City string
	State string
	zip string
	Latitude float32 `gorm:"type:decimal;precision:11,8"`
	Longitude float32 `gorm:"type:decimal;precision:11,8"`
}

// 默认表名有问题, 设置表名
func (Cafe) TableName() string {
	return "cafes"
}


// 查询数据库中所有咖啡店
func GetAllCafes() Cafe {
	var cafes Cafe

	Db.First(&cafes)

	fmt.Printf("cafes %v", cafes)

	return cafes
}
package models

import (
	// "fmt"
	// "github.com/jinzhu/gorm"
	. "roast-server/database"
	"time"
)

type Cafe struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name" binding:"required,max=50"`
	Address   string     `json:"address" binding:"required"`
	City      string     `json:"city" binding:"required"`
	State     string     `json:"state" binding:"required"`
	Zip       string     `json:"zip" binding:"required"`
	Latitude  string    `gorm:"type:varchar(255)"  json:"latitude"`
	Longitude string    `gorm:"type:varchar(255)"  json:"longtitude"`
}

// 默认表名有问题, 设置表名
func (Cafe) TableName() string {
	return "cafes"
}

// 查询数据库中所有咖啡店
func GetAllCafes() []Cafe {
	var cafes []Cafe

	Db.Find(&cafes)

	// fmt.Printf("cafes %v", cafes)

	return cafes
}

// 获取咖啡店详情
func GetCafe(id string) Cafe {
	var cafe Cafe

	Db.First(&cafe, id)

	return cafe
}

// 创建新的咖啡店
func PostNewCafe(cafe *Cafe) {
	Db.Create(&cafe)
}
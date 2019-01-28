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
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	City      string     `json:"city"`
	State     string     `json:"state"`
	Zip       string     `json:"zip"`
	Latitude  float32    `gorm:"type:decimal;precision:11,8"  json:"latitude"`
	Longitude float32    `gorm:"type:decimal;precision:11,8"  json:"longtitude"`
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
func GetCafe(id int) Cafe {
	var cafe Cafe

	Db.First(&cafe, id)

	return cafe
}

// 创建新的咖啡店
func PostNewCafe(cafe *Cafe) {
	Db.Create(&cafe)
}
package models

import (
	// "fmt"
	// "github.com/jinzhu/gorm"
	. "roast-server/database"
	"time"
)


type BrewMethod struct {
	ID        uint       `gorm:"primary_key" json:"id"`	
	Method    string  `json:"method"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Cafes []*Cafe `gorm:"many2many:cafes_brew_methods;" json:"cafes"`
	CafesCount int `gorm:"-" json:"cafes_count"`
}

// 获取所有冲泡方法
func GetBrewMethods() []BrewMethod {
	var brewMethods []BrewMethod

	Db.Preload("Cafes").Find(&brewMethods)

	// 计算冲泡方法对应的咖啡店数量
	var length int
	for k, v := range brewMethods {
		length = len(v.Cafes)
		brewMethods[k].CafesCount = length
	}

	return brewMethods
}
package models

import (
	// "fmt"
	// "github.com/jinzhu/gorm"
	. "roast-server/database"
	"time"
)

type Cafe struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name" binding:"required,max=50"`
	Address   string     `json:"address" binding:"required"`
	City      string     `json:"city" binding:"required"`
	State     string     `json:"state" binding:"required"`
	Zip       string     `json:"zip" binding:"required"`
	ParentId    uint        `gorm:"type:int;default:0"  json:"parent_id"`
	LocationName string  `gorm:"type:varchar(255)"  json:"location_name"`
	Roaster int  		`gorm:"type:tinyint(1)"  json:"roaster"`
	Website  string 	`gorm:"type:varchar(255)"  json:"website"`
	Description  string `gorm:"type:varchar(255)"  json:"description"`
	AddedBy   int        `gorm:"type:int;default:0"  json:"added_by"`
	Latitude  string    `gorm:"type:varchar(255)"  json:"latitude"`
	Longitude string    `gorm:"type:varchar(255)"  json:"longitude"`
	BrewMethods []*BrewMethod `gorm:"many2many:cafes_brew_methods;association_autoupdate:false" json:"brew_methods"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Children  []Cafe `gorm:"foreignkey:ParentId" json:"children"`
	// Parent Cafe `gorm:"foreignkey:ID;association_foreignkey:ParentId" json:"parent"`
	Likes []User `gorm:"many2many:users_cafes_likes"`
	UserLike bool `gorm:"-" json:"user_like"`
	Tags []Tag `gorm:"many2many:cafes_users_tags;association_jointable_foreignkey:cafe_id;jointable_foreignkey:tag_id;"`
}

// 默认表名有问题, 设置表名
func (Cafe) TableName() string {
	return "cafes"
}

// 查询数据库中所有咖啡店
func GetAllCafes() []Cafe {
	var cafes []Cafe

	Db.Preload("BrewMethods").Find(&cafes)

	// fmt.Printf("cafes %v", cafes)

	return cafes
}

// 获取咖啡店详情
func GetCafe(id string) Cafe {
	var cafe Cafe

	Db.Preload("BrewMethods").Preload("Tags").First(&cafe, id)

	var likes []UsersCafesLike

	Db.Where("user_id = ? and cafe_id = ?", 2, cafe.ID).Find(&likes)

	if len(likes) > 0 {
		cafe.UserLike = true
	}

	return cafe
}

// 创建新的咖啡店
func PostNewCafe(cafe *Cafe) {
	Db.Create(&cafe)
}

// 喜欢咖啡店
func PostLikeCafe(id string) bool {
	var cafe Cafe

	Db.First(&cafe, id)

	Db.Model(&cafe).Association("Likes").Append(User{ID: 2})

	return true
}

// 取消喜欢咖啡店
func DeleteLikeCafe(id string) bool {
	var cafe Cafe

	Db.First(&cafe, id)

	Db.Model(&cafe).Association("Likes").Delete(User{ID: 2})

	return true
}
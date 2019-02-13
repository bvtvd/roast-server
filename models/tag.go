package models

import (
	"fmt"
	. "roast-server/database"
)

type Tag struct {
	ID        uint       `gorm:"primary_key" json:"id"`	
	Name    string  `gorm:"UNIQUE_INDEX;type:varchar(255)" json:"name"`
	Cafes []Cafe `gorm:"many2many:cafes_users_tags;association_jointable_foreignkey:tag_id;jointable_foreignkey:cafe_id;"`
}


func GetTags(name string) []Tag {
	var tags []Tag

	query := Db

	if name != "" {
		t, _ := fmt.Printf("%%s%", name)
		query = query.Where("name like ?", t)
	}

	query.Find(&tags)

	return tags
}
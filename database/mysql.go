package database

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    // . "roast-server/models"	// 这里的代码会造成包的循环导入
    "os"
)

var Db *gorm.DB

func InitDb() {
	var err error

	// Db, err = gorm.Open("mysql", "root:root@tcp(192.168.1.193:3306)/roast?charset=utf8&parseTime=True&loc=Local")
	Db, err = gorm.Open("mysql", os.Getenv("MYSQL"))

	if err != nil {
		panic("database connetc error: " + err.Error())
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	Db.LogMode(true)


	// 自动迁移模式
	// Db.AutoMigrate(
	// 	&Cafe{},
	// )

}
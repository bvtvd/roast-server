package main  

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    . "roast-server/models"	// 这里的代码会造成包的循环导入
)

func main() {
	Db, err := gorm.Open("mysql", "root:root@tcp(192.168.1.193:3306)/roast?charset=utf8&parseTime=True&loc=Local")
	defer Db.Close()

	if err != nil {
		panic("database connetc error: " + err.Error())
	}

	//自动迁移模式
	Db.AutoMigrate(
		&Cafe{},
	)
}
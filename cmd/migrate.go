package main  

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    . "roast-server/models"	// 这里的代码会造成包的循环导入
    "fmt"
)

func main() {
	Db, err := gorm.Open("mysql", "root:root@tcp(192.168.43.245:3306)/roast?charset=utf8&parseTime=True&loc=Local")
	// defer Db.Close()

	if err != nil {
		panic("database connetc error: " + err.Error())
	}

	Db.LogMode(true)

	if Db.HasTable("cafes") {
		fmt.Println("got table")
	}else{
		fmt.Println("no table")
	}

	//自动迁移模式
	Db.AutoMigrate(
		&Cafe{}, &BrewMethod{}, &CafesBrewMethod{},
	)
	fmt.Println("migrate finished")
}
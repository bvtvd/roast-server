package main  

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    . "roast-server/models"	// 这里的代码会造成包的循环导入
    "fmt"
    "github.com/joho/godotenv"
	"os"
)

func main() {
	// 读取配置文件
	if err := godotenv.Load(".env"); err != nil {
		panic("config file read error")
	}

	Db, err := gorm.Open("mysql", os.Getenv("MYSQL"))
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
		&Cafe{}, &BrewMethod{}, &CafesBrewMethod{}, &UsersCafesLike{}, &User{}, &Tag{}, &CafesUsersTag{},
	)
	fmt.Println("migrate finished")
}
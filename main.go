package main

import (
	. "roast-server/database"
	"roast-server/routers"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// 读取配置文件
	if err := godotenv.Load(".env"); err != nil {
		panic("config file read error")
	}

	// defer Db.Close()
	InitDb()

	routers := routers.InitRouter()

	routers.Run(":" + os.Getenv("PORT"))
}

package main

import (
	. "roast-server/database"
	"roast-server/routers"
)

func main() {
	// defer Db.Close()
	InitDb()

	routers := routers.InitRouter()

	routers.Run(":9000")
}

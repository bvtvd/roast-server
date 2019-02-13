package models

// import (
// 	// "fmt"
// 	// "github.com/jinzhu/gorm"
// 	// . "roast-server/database"
// 	//"time"
// )


type UsersCafesLike struct {
	UserId        uint       `json:"user_id"`	
	CafeId        uint       `json:"cafe_id"`	
}
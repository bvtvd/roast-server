package controllers

import (
	"github.com/gin-gonic/gin"
	. "roast-server/models"
	"roast-server/utils"
	"fmt"
	// "strconv"
	// validator "gopkg.in/go-playground/validator.v9"
)

// type CommonError struct {
//     Errors map[string]interface{} `json:"errors"`
// }

// func NewValidatorError(err error) CommonError {
//     res := CommonError{}
//     res.Errors = make(map[string]interface{})

//     errs := err.(validator.ValidationErrors)

//     for _, e := range errs {
//         res.Errors[e.Field()] = e.ActualTag()
//     }
//     return res
// }

// 获取所有咖啡店
func CafesGetCafes(c *gin.Context) {

	cafes := GetAllCafes()
	// fmt.Printf("cafes %v", cafes)

	c.JSON(200, cafes)
	// c.String(200, "hello")
}

// 创建新的咖啡店
func CafesPostNewCafe(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	city := c.PostForm("city")
	state := c.PostForm("state")
	zip := c.PostForm("zip")

	cafe := Cafe{Name: name, Address: address, City: city, State: state, Zip: zip}
	fmt.Printf("[CAFE]: %v \n", cafe)
	// 数据验证
	if err := c.ShouldBind(&cafe); err != nil {
		c.JSON(422, err.Error())
		return
	}

	coordinate, _ := utils.GeocodeAddress(cafe.Address, cafe.City, cafe.State)

	cafe.Latitude = coordinate.Lat
	cafe.Longitude = coordinate.Lng
	cafe.DeletedAt = nil
	fmt.Printf("[CAFE]: %v \n", cafe)
	PostNewCafe(&cafe)

	c.JSON(201, cafe)
}

// 获取咖啡店详情
func CafesGetCafe(c *gin.Context) {
	id := c.Param("id")

	cafe := GetCafe(id)

	c.JSON(200, cafe)
}


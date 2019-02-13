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

// 定义创建新咖啡店要接收的数据
type Location struct {
	Name string 	`form:"name" json:"name" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
	City string `form:"city" json:"city" binding:"required"`
	State string  `form:"state" json:"state" binding:"required"`
	Zip string  `form:"zip" json:"zip"`
	MethodsAvailable []int   `form:"methodsAvailable" json:"methodsAvailable" binding:"required"`
}

type NewCafe struct {
	Name string `form:"name" json:"name" binding:"required"`
	Website string `form:"website" json:"website" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Roaster bool `form:"roaster" json:"roaster"`
	Locations []Location `form:"locations" json:"locations" binding:"dive"`
}


// 创建新的咖啡店
func CafesPostNewCafe(c *gin.Context) {
	// name := c.PostForm("name")
	// address := c.PostForm("address")
	// city := c.PostForm("city")
	// state := c.PostForm("state")
	// zip := c.PostForm("zip")

	// cafe := Cafe{Name: name, Address: address, City: city, State: state, Zip: zip}
	// fmt.Printf("[CAFE]: %v \n", cafe)
	// // 数据验证
	// if err := c.ShouldBind(&cafe); err != nil {
	// 	c.JSON(422, err.Error())
	// 	return
	// }

	// coordinate, _ := utils.GeocodeAddress(cafe.Address, cafe.City, cafe.State)

	// cafe.Latitude = coordinate.Lat
	// cafe.Longitude = coordinate.Lng
	// cafe.DeletedAt = nil
	// fmt.Printf("[CAFE]: %v \n", cafe)
	// PostNewCafe(&cafe)

	// c.JSON(201, cafe)

	
	var createParam NewCafe
	// c.BindJSON(&new)
	// c.ShouldBind(&createParam)
	if err := c.ShouldBind(&createParam); err != nil {
		c.JSON(422, err.Error())
		return
	}
	// 先绑定数据进行数据验证
	// var cafes []Cafe
	fmt.Println("datas==========================================")
	fmt.Printf("%v \n", createParam)
	// 将提交的数据重组
	var cafes []Cafe
	var brewMethods []*BrewMethod
	for _, v := range createParam.Locations {
		coordinate, _ := utils.GeocodeAddress(v.Address, v.City, v.State)
		brewMethods = make([]*BrewMethod, 0)

		for _, m := range v.MethodsAvailable {
			brewMethods = append(brewMethods, &BrewMethod{ID: uint(m)})
		}

		var r int 
		if createParam.Roaster {
			r = 1
		}else{
			r = 0
		}

		cafes = append(cafes, Cafe{
			Name: createParam.Name, 
			LocationName: v.Name, 
			Address: v.Address, 
			City: v.City, 
			State: v.State, 
			Zip: v.Zip,
			Latitude: coordinate.Lat,
			Longitude: coordinate.Lng,
			Roaster: r,
			Website: createParam.Website,
			Description: createParam.Description,
			AddedBy: 1,
			BrewMethods: brewMethods})
	}

	fmt.Println("cafes==========================================")
	fmt.Printf("%v \n", cafes)

	// 数据入库
	// var parent Cafe
	for k, _ := range cafes {
		if k > 0 {
			cafes[k].ParentId = cafes[0].ID
		}
		PostNewCafe(&cafes[k])
	}

	c.JSON(201, gin.H{"data": cafes})
}

// 获取咖啡店详情
func CafesGetCafe(c *gin.Context) {
	id := c.Param("id")

	cafe := GetCafe(id)

	c.JSON(200, cafe)
}


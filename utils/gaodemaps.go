package utils

import (
	"net/http"
	"fmt"
	"net/url"
	"encoding/json"
	"strings"
)



const (
	API_KEY string = "f2dba2d71fa9d7771ffb488b2b22023a"
)

// 经纬度
type Coordinate struct {
	Lat string
	Lng string
}


type GeoSearchResult struct {
	Status string `json:"status"`
	Geocodes []*Geocode `json:"geocodes"`
}

type Geocode struct {
	Location string `json:"location"`
}

// 根据地址返回经纬度信息
func GeocodeAddress(address string, city string, state string) (Coordinate, error) {
	var coordinate Coordinate

	v := url.Values{}

	v.Add("key", API_KEY)
	v.Add("address", state + city + address)

	body:= v.Encode()

	fmt.Printf("[ENCODE BODY]: %v \n", body)

	url := "https://restapi.amap.com/v3/geocode/geo?" + body

	response, err := http.Get(url)

	defer response.Body.Close()

	if err != nil{
		return coordinate, err
	}

	var result GeoSearchResult
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Printf("[DECODE ERR]: %v \n", err)
		return coordinate, err
	}

	fmt.Printf("[DECODE RESULT]: %v \n", result.Geocodes[0])

	// TODO:: 返回地理编码
	fmt.Printf("[LENGTH]: %d \n", len(result.Geocodes))

	if len(result.Geocodes) > 0 {

		location := result.Geocodes[0].Location
		fmt.Printf("[LOCATION]: %v \n", location)

		locationArr := strings.Split(string(location), ",")
		fmt.Printf("[LOCATION Arr]: %v \n", locationArr)
		
		coordinate = Coordinate{locationArr[1], locationArr[0]}
	}
	
	return coordinate, nil
}
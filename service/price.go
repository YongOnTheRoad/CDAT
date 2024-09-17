package service

import (
	"fmt"
	"regexp"
	"strings"

	"cdat1.0/utils"
	"github.com/gin-gonic/gin"
)

func TestUrl(c *gin.Context) {
	c.JSON(200, "success")
}

func GetPrice(c *gin.Context) {
	var paths []string
	var prices []utils.HouseOneDay
	csvpaths := utils.GetAllFile("C://work//go_file", paths)
	for i := 0; i < len(csvpaths); i++ {
		/*fmt.Println("file path is:", csvpaths[i])*/
		splitResult := strings.Split(csvpaths[i], "//")
		date := splitResult[len(splitResult)-1]
		/*fmt.Println("data is:", date)*/
		regularExpression, err := regexp.Compile("[0-9]+")
		if err != nil {
			fmt.Println("regular expression failed")
		}
		matches := regularExpression.FindAllString(date, -1)
		/*fmt.Println("match result is:", matches)*/
		var currentDate string
		for _, i := range matches {
			currentDate += i
		}

		price := utils.ReadCsvFile(csvpaths[i], currentDate[:8])
		prices = append(prices, price)
	}

	c.JSON(200, prices)
}

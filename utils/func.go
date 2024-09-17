// utils/func.go

package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

type HouseInfo struct {
	RealEstate string `json:"realEstate"`
	TotalPrice string `json:"totalPrice"`
	UnitPrice  string `json:"unitPrice"`
	Flood      string `json:"flood"`
	Summary    string `json:"summary"`
	Heat       string `json:"heat"`
}

type HouseOneDay struct {
	HouseDay    string      `json:"houseDay"`
	Information []HouseInfo `json:"information"`
}

func TestFunc(c *gin.Context) {
	klog.Infoln("Hello Go Modules!")
	c.JSON(200, 100)
}

func ReadCsvFile(filePath string, date string) HouseOneDay {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var tmpDayInfo = HouseOneDay{HouseDay: date}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read record failed:", err)
		}
		/* fmt.Println("read result: ", record) */
		infoHouse := HouseInfo{record[0], record[1], record[2], record[3], record[4], record[5]}
		/* fmt.Printf(infoHouse.Flood) */
		tmpDayInfo.Information = append(tmpDayInfo.Information, infoHouse)
	}
	/* fmt.Println("read result: ", tmpDayInfo) */
	return tmpDayInfo
}

func GetAllFile(dirpath string, filepaths []string) []string {
	reader, err := ioutil.ReadDir(dirpath)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}
	for _, file := range reader {
		if !file.IsDir() {
			filepath := dirpath + "//" + file.Name()
			/* fmt.Println("file path is:", filepath) */
			filepaths = append(filepaths, filepath)
		}
	}
	return filepaths
}

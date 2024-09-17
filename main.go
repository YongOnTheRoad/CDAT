// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"cdat1.0/cmd"
	/*"cdat1.0/utils"*/)

func main() {
	name := "custom data analysis tool"
	fmt.Println("wherecome to", name)

	/*oneDayInfo := utils.ReadCsvFile("C://work//go_file//house_price_hong_shan_20240131231051.csv")
	fmt.Println(oneDayInfo)
	var csvpaths []string
	csvpaths, _ = utils.GetAllFile("C://work//go_file", csvpaths)
	fmt.Println(csvpaths)*/

	app := &cli.App{
		Name:  "cdat",
		Usage: "me",
		//Action: func(c *cli.Context) error {
		//	fmt.Println("boom! I say!")
		//	return nil
		//},
		Commands: []*cli.Command{
			cmd.Web,
		},
		DefaultCommand: "web",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

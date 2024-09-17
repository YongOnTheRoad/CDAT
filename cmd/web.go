package cmd

import (
	"fmt"
	"log"

	"cdat1.0/service"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:   "web",
	Usage:  "start web server",
	Action: runWeb,
}

func runWeb(c *cli.Context) error {
	r := gin.Default()
	v1 := r.Group("/api/cdat")
	v1.GET("/prices", service.GetPrice)
	v1.GET("/test", service.TestUrl)
	err := r.Run("0.0.0.0:8001")
	if err != nil {
		log.Fatal("start web server failed")
	}
	fmt.Println("start web server success!")
	return err
}

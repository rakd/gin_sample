package main

import (
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/rakd/gin_sample/app/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	router := gin.Default()

	//gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/robots.txt", "./assets/robots.txt")
	router.Use(gin.Recovery())
	router.GET("/", controllers.AppIndex)
	router.Run(":3000")
}

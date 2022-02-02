package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "assets")

	router.GET("/", homeRoute)

	router.Run()
}

func homeRoute(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFS("/resources", http.Dir("resources"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/h/:hash", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"hash":    c.Param("hash"),
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:80
}

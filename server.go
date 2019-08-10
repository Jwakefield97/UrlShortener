package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))[0:8]
}

func main() {
	var urls sync.Map
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFS("/resources", http.Dir("resources"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/new", func(c *gin.Context) {
		newURL := c.PostForm("url")
		urlHash := hash(newURL)
		urls.Store(urlHash, newURL)
		c.JSON(200, gin.H{
			"hash": urlHash,
		})
	})

	r.GET("/h/:hash", func(c *gin.Context) {
		url, ok := urls.Load(c.Param("hash"))
		if ok {
			c.Redirect(http.StatusFound, url.(string))
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.Run(":80") // listen and serve on 0.0.0.0:80
}

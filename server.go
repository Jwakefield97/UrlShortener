package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))[0:8]
}

func main() {
	shortenLink := "http://localhost/h/"
	urls := map[string]string{}
	urls["1"] = "https://www.google.com/search?q=images&rlz=1C1CHBF_enUS796US796&tbm=isch&source=iu&ictx=1&fir=8c_UAo3gH_220M%253A%252CShwNVOdFBcmkxM%252C_&vet=1&usg=AI4_-kQpJbM15TX5VhVDvz1rgddeX4UTow&sa=X&ved=2ahUKEwj8sIK0lfnjAhUWCs0KHU6MBUsQ9QEwAXoECAYQBg#imgrc=8c_UAo3gH_220M:"
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFS("/resources", http.Dir("resources"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/search", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.html", nil)
	})

	//return the map of currently mapped hashs to urls
	r.GET("/urls", func(c *gin.Context) {
		c.JSON(200, urls)
	})

	//add a new hash to the url
	r.POST("/new", func(c *gin.Context) {
		newURL := c.PostForm("url")
		urlHash := hash(newURL)
		urls[urlHash] = newURL
		c.JSON(200, gin.H{
			"url": shortenLink + urlHash,
		})
	})

	//redirect to mapped url
	r.GET("/h/:hash", func(c *gin.Context) {
		url, ok := urls[c.Param("hash")]
		if ok {
			c.Redirect(http.StatusFound, url)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.Run(":80") // listen and serve on 0.0.0.0:80
}

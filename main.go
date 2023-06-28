package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pmylund/go-cache"
)

func main() {
	clientCache := cache.New(5*time.Minute, 10*time.Minute)

	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {

		_, found := clientCache.Get("api")
		if found {
			c.JSON(http.StatusOK, gin.H{
				"status": "Cached.",
			})
			return
		}

		clientCache.Set("api", true, cache.NoExpiration)
		c.JSON(http.StatusOK, gin.H{
			"status": "First.",
		})
	})
	r.Run(":1111") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

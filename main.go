package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/feed", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"feed":    gin.H{},
		})
	})

	r.Run(":8080")
}

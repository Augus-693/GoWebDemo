package main

import "github.com/gin-gonic/gin"

/**
 * @Project GoWebDemo
 * @File    book.go
 * @Author  Augus Lee
 * @Date    2022/11/21 14:49
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run()
}

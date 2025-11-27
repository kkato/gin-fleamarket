package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkato/gin-fleamarket/models"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/items", func(c *gin.Context) {
		c.JSON(200, items)
	})
	router.Run("localhost:8080") // listens on 0.0.0.0:8080 by default
}

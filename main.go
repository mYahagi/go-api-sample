package main

import (
	"./src/app/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		repo := repository.FishRepository{}
		fish := repo.FindById(1)

		c.JSON(200, gin.H{
			"message": fish.NAME,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

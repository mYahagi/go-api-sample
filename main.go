package main

import (
	"math/rand"

	"./src/app/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		repo := repository.FishRepository{}

		count := repo.Count()

		fish := repo.FindById(randomInt(count))

		c.JSON(200, gin.H{
			"釣れた魚": fish.NAME,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func randomInt(max int) int {
	value := 0
	for {
		value = rand.Intn(max)
		if value > 0 {
			break
		}
	}

	return value
}

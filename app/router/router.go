package router

import (
	"../../app/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func Router() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		repo := repository.NewFishRepository()

		count := repo.Count()
		fish := repo.FindById(randomInt(count))

		c.JSON(200, gin.H{
			"釣れた魚": fish.NAME,
		})
	})
	r.Run()
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

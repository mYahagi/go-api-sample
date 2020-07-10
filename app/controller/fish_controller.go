package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"

	"../infrastructure/repository"
)

func Catch(c *gin.Context) {
	repo := repository.NewFishRepository()

	count, err := repo.Count()
	if err != nil {
		c.String(http.StatusBadRequest, "釣れませんでした")
	}

	fish, err := repo.FindById(randomInt(*count))
	if err != nil {
		c.String(http.StatusBadRequest, "釣れませんでした")
	}

	c.JSON(200, gin.H{
		"釣れた魚": fish.NAME,
	})
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

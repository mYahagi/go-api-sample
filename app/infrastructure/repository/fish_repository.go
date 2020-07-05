package repository

import (
	"../../domain"
	"../database"
)

type FishRepository struct{}

func (repo *FishRepository) FindById(identifier int) (fish domain.Fish) {
	db := database.GormConnect()
	defer db.Close()

	db.First(&fish, identifier)

	return fish
}

func (repo *FishRepository) Count() int {
	db := database.GormConnect()
	defer db.Close()

	var fishes domain.Fishes
	count := 0
	db.Find(&fishes).Count(&count)

	return count
}

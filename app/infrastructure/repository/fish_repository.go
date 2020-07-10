package repository

import (
	"../database"
	"../dto"
	"github.com/jinzhu/gorm"
)

type FishRepository struct {
	db *gorm.DB
}

func (repo *FishRepository) FindById(identifier int) (fish dto.Fish) {
	repo.db.First(&fish, identifier)

	return fish
}

func (repo *FishRepository) Count() int {
	var fishes dto.Fishes
	count := 0

	repo.db.Find(&fishes).Count(&count)

	return count
}

func NewFishRepository() *FishRepository {
	fishRepository := new(FishRepository)
	fishRepository.db = database.Connect()

	return fishRepository
}

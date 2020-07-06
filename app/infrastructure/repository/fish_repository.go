package repository

import (
	"../../domain"
	"../database"
	"github.com/jinzhu/gorm"
)

type FishRepository struct {
	db *gorm.DB
}

func (repo *FishRepository) FindById(identifier int) (fish domain.Fish) {
	repo.db.First(&fish, identifier)
	defer repo.db.Close()

	return fish
}

func (repo *FishRepository) Count() int {
	var fishes domain.Fishes
	count := 0

	repo.db.Find(&fishes).Count(&count)
	defer repo.db.Close()

	return count
}

func NewFishRepository() *FishRepository {
	fishRepository := new(FishRepository)
	fishRepository.db = database.Connect()

	return fishRepository
}

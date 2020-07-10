package repository

import (
	"../database"
	"../dto"
	"github.com/jinzhu/gorm"
)

type FishRepository struct {
	db *gorm.DB
}

func (repo *FishRepository) FindById(identifier int) (*dto.Fish, error) {
	var fish dto.Fish
	if err := repo.db.First(&fish, identifier).Error; err != nil {
		return nil, err
	}

	return &fish, nil
}

func (repo *FishRepository) Count() (*int, error) {
	var fishes dto.Fishes
	var count int

	if err := repo.db.Find(&fishes).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

func NewFishRepository() *FishRepository {
	fishRepository := new(FishRepository)
	fishRepository.db = database.Connect()

	return fishRepository
}

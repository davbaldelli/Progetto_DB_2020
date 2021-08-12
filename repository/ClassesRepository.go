package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type ClassesRepository struct {
	Db *gorm.DB
}

func (c ClassesRepository) GetAllClasses() ([]models.CarClass, error) {
	var dbClasses []CarClass

	if err := c.Db.Find(&dbClasses).Error; err != nil {
		return nil, err
	}

	var classes []models.CarClass
	for _, class := range dbClasses {
		classes = append(classes, models.CarClass{Name: class.Name})
	}
	return classes, nil
}

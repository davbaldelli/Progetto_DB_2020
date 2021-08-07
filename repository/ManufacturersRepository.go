package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type ManufacturerRepository struct {
	Db *gorm.DB
}

func (m ManufacturerRepository) GetAllManufacturers() ([]models.Brand, error) {
	var dbManufacturers []Manufacturer
	if err := m.Db.Find(&dbManufacturers).Error; err != nil {
		return nil, err
	}

	var manufacturers []models.Brand

	for _, manufacturer := range dbManufacturers {
		manufacturers = append(manufacturers, models.Brand{
			Name:   manufacturer.Name,
			Nation: manufacturer.Nation,
		})
	}

	return manufacturers, nil
}

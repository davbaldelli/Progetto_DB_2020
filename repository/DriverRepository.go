package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type DriverRepository struct {
	Db *gorm.DB
}

func (d DriverRepository) GetAllDrivers() ([]models.Driver, error) {
	var dbDrivers []Driver
	if err := d.Db.Find(&dbDrivers).Error; err != nil {
		return nil, err
	}

	var drivers []models.Driver

	for _, driver := range dbDrivers {
		drivers = append(drivers, models.Driver{
			Name:      driver.Name,
			Surname:   driver.Surname,
			CF:        driver.Cf,
			Sex:       driver.Sex,
			Birthdate: driver.Birthdate,
			Nation:    driver.Nation,
		})
	}

	return drivers, nil
}

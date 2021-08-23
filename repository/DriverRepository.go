package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type DriverRepository struct {
	Db *gorm.DB
}

func dbDriversToModels(dbDrivers []Driver) []models.Driver {

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

	return drivers
}

func (d DriverRepository) GetAllDrivers() ([]models.Driver, error) {
	var dbDrivers []Driver
	if err := d.Db.Find(&dbDrivers).Error; err != nil {
		return nil, err
	}

	return dbDriversToModels(dbDrivers), nil
}

func (d DriverRepository) GetFiveDriversWithMoreRaces() ([]models.DriverRaces, error) {

	var drivers []models.DriverRaces
	if err := d.Db.Table("drivers").
		Select("drivers.*", "COUNT(races.id) AS `races`").
		Joins("LEFT JOIN driver_entries ON drivers.cf = driver_entries.driver").
		Joins("LEFT JOIN entries ON driver_entries.entry = entries.id").
		Joins("LEFT JOIN championships ON entries.championship = championships.id").
		Joins("LEFT JOIN races ON races.championship = championships.id").
		Group("drivers.cf").
		Order("`races` DESC").
		Limit(10).
		Find(&drivers).Error; err != nil {
		return nil, err
	}

	return drivers, nil
}

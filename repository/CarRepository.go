package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type CarRepository struct {
	Db *gorm.DB
}

func dbCarsToModels(dbCars []Car) []models.Car {
	var cars []models.Car
	for _, car := range dbCars {
		cars = append(cars, models.Car{
			Model:        car.Model,
			Year:         car.Year,
			Brand:        models.Brand{Name: car.Brand},
			Class:        car.Class,
			Drivetrain:   car.Drivetrain,
			Transmission: car.Transmission,
		})
	}
	return cars
}

func (c CarRepository) GetAllCars() ([]models.Car, error) {
	var dbCars []Car

	if err := c.Db.Find(&dbCars).Error; err != nil {
		return nil, err
	}

	return dbCarsToModels(dbCars), nil
}

func (c CarRepository) GetChampionshipCars(championship models.Championship) ([]models.Car, error) {
	var dbCars []Car

	if err := c.Db.
		Joins("JOIN entries ON cars.id = entries.car").
		Joins("JOIN championships ON championships.id = entries.championship").
		Where("championships.name = ? AND championships.year = ?", championship.Name, championship.Year).
		Distinct().Find(&dbCars).Error; err != nil {
		return nil, err
	}

	return dbCarsToModels(dbCars), nil
}

func (c CarRepository) GetDriverCarOnCircuit(driver models.Driver, track models.Track) ([]models.Car, error) {
	var dbCars []Car

	if err := c.Db.Distinct().
		Joins("JOIN entries ON entries.car = cars.id").
		Joins("JOIN driver_entries ON entries.id = driver_entries.entry").
		Joins("JOIN championships ON championships.id = entries.championship").
		Joins("JOIN races ON races.championship = championships.id").
		Joins("JOIN layouts ON races.layout = layouts.id").
		Where("driver_entries.driver = ?", driver.CF).
		Where("layouts.track = ?", track.Name).
		Find(&dbCars).
		Error; err != nil {
		return nil, err
	}
	return dbCarsToModels(dbCars), nil
}

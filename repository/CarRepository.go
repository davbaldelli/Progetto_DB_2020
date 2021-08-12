package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type CarRepository struct {
	Db *gorm.DB
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

	return cars, nil
}

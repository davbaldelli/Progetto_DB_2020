package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type StatisticsRepository struct {
	Db *gorm.DB
}

func (s StatisticsRepository) GetBrandCarsUsage(brandName string) ([]models.CarUsage, error) {
	var statistics []models.CarUsage
	//return list of cars of a specified brand with the number of total entries (across al championships)
	if err := s.Db.Table("entries").
		Select("cars.*, count(entries.id) as 'usage'").
		Joins("right join cars ON entries.car = cars.id").
		Group("cars.id").Having("brand = ?", brandName).
		Scan(&statistics).Error; err != nil {
		return nil, err
	}
	return statistics, nil
}

func (s StatisticsRepository) GetTrackLayoutsUsage(trackName string) ([]models.LayoutUsage, error) {
	var statistics []models.LayoutUsage
	//returns the number of races of all layout of a specified track
	if err := s.Db.Table("races").
		Select("layouts.*", "count(races.id) as 'usage'").
		Joins("right join layouts on races.layout = layouts.id").
		Group("layouts.id").Having("track = ?", trackName).
		Scan(&statistics).Error; err != nil {
		return nil, err
	}
	return statistics, nil
}

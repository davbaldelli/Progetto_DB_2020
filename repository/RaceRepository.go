package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type RacesRepository struct {
	Db *gorm.DB
}

func dbRaceToEntity(dbRace Race) models.Race {
	return models.Race{
		Name:             dbRace.Name,
		Date:             dbRace.Datetime,
		Track:            models.Track{Name: dbRace.Track},
		LayoutName:       dbRace.Layout,
		ChampionshipName: dbRace.Championship,
	}
}

func (r RacesRepository) GetChampionshipRaces(championship models.Championship) ([]models.Race, error) {
	var dbRaces []Race

	if err := r.Db.
		Select("races.name, races.datetime", "tracks.name AS track", "layouts.name AS layout").
		Where("championships.name = ? AND championships.year = ?", championship.Name, championship.Year).
		Joins("join championships ON championships.id = races.championship").
		Joins("join layouts ON layouts.id = races.layout").
		Joins("join tracks ON layouts.track = tracks.name").
		Find(&dbRaces).Error; err != nil {
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil

}

func (r RacesRepository) GetRacesByClass(class string) ([]models.Race, error) {
	var dbRaces []Race
	if err := r.Db.
		Table("championship_classes").
		Select("races.name, races.datetime", "championships.name AS championship", "tracks.name AS track", "layouts.name AS layout").
		Where("class = ?", class).
		Joins("join races ON races.championship = championship_classes.championship").
		Joins("join championships ON championships.id = races.championship").
		Joins("join layouts ON layouts.id = races.layout").
		Joins("join tracks ON layouts.track = tracks.name").
		Find(&dbRaces).Error; err != nil {
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil
}

func (r RacesRepository) GetRacesByTeam(teamName string) ([]models.Race, error) {
	var dbRaces []Race
	if err := r.Db.Table("entries").Distinct().
		Where("team = ?", teamName).
		Select("races.name, races.datetime", "championships.name AS championship", "tracks.name AS track", "layouts.name AS layout").
		Joins("join championships on championships.id = entries.championship").
		Joins("join races on races.championship = championships.id").
		Joins("join layouts ON layouts.id = races.layout").
		Joins("join tracks ON layouts.track = tracks.name").
		Find(&dbRaces).Error; err != nil {
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil
}

func (r RacesRepository) GetDriversRacesByNationality(nation string) ([]models.Race, error) {
	var dbRaces []Race
	if err := r.Db.Table("drivers").Distinct().
		Where("drivers.nation = ?", nation).
		Select("races.name, races.datetime", "championships.name AS championship", "tracks.name AS track", "layouts.name AS layout").
		Joins("join driver_entries on drivers.cf = driver_entries.driver").
		Joins("join entries on  driver_entries.entry = entries.id").
		Joins("join championships on championships.id = entries.championship").
		Joins("join races on races.championship = championships.id").
		Joins("join layouts ON layouts.id = races.layout").
		Joins("join tracks ON layouts.track = tracks.name").
		Find(&dbRaces).Error; err != nil {
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil
}

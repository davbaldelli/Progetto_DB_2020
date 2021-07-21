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

func (r RacesRepository) GetIncomingRacesByClass(class string) ([]models.Race, error) {
	var dbRaces []Race
	if err := r.Db.
		Table("championship_classes").
		Select("races.*", "layouts.track AS track_name").
		Where("class = ?", class).
		Joins("join races ON races.championship = championship_classes.championship").
		Joins("join championships ON championships.id = races.championship").
		Joins("join layouts ON layouts.id = races.layout").
		Find(&dbRaces).Error; err != nil{
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil
}

func (r RacesRepository) GetIncomingRacesByTeam(teamName string) ([]models.Race, error) {
	var dbRaces []Race
	if err := r.Db.Table("entries").Distinct().
		Where("team = ?", teamName).
		Select("races.*",  "layouts.track AS track_name").
		Joins("join championships on championships.id = entries.championship").
		Joins("join races on races.championship = championships.id").
		Joins("join layouts ON layouts.id = races.layout").
		Find(&dbRaces).Error ; err != nil{
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
		Select("races.*",  "layouts.track AS track_name").
		Joins("join driver_entries on drivers.cf = driver_entries.driver").
		Joins("join entries on  driver_entries.entry = entries.id").
		Joins("join championships on championships.id = entries.championship").
		Joins("join races on races.championship = championships.id").
		Joins("join layouts ON layouts.id = races.layout").

		Find(&dbRaces).Error ; err != nil{
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil
}


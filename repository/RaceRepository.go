package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type RacesRepository struct {
	Db *gorm.DB
}

type raceQuery func(*[]Race) *gorm.DB

func (r RacesRepository) selectRacesWithQuery(query raceQuery) ([]models.Race, error) {
	var dbRaces []Race

	if err := query(&dbRaces).Error; err != nil {
		return nil, err
	}

	var races []models.Race

	for _, dbRace := range dbRaces {
		races = append(races, dbRaceToEntity(dbRace))
	}
	return races, nil
}

func dbRaceToEntity(dbRace Race) models.Race {
	return models.Race{
		Name:         dbRace.Name,
		Date:         dbRace.Datetime,
		Track:        models.Track{Name: dbRace.Track, Nation: dbRace.TrackNation, Location: dbRace.TrackLocation},
		LayoutName:   dbRace.Layout,
		Championship: models.Championship{Name: dbRace.ChampionshipName, Year: dbRace.ChampionshipYear},
	}
}

func (r RacesRepository) GetChampionshipRaces(championship models.Championship) ([]models.Race, error) {
	return r.selectRacesWithQuery(func(dbRaces *[]Race) *gorm.DB {
		return r.Db.
			Select("races.name, races.datetime", "championships.name AS championship_name", "championships.year AS championship_year", "tracks.name AS track", "layouts.name AS layout", "tracks.nation as track_nation", "tracks.location as track_location").
			Where("championships.name = ? AND championships.year = ?", championship.Name, championship.Year).
			Joins("join championships ON championships.id = races.championship").
			Joins("join layouts ON layouts.id = races.layout").
			Joins("join tracks ON layouts.track = tracks.name").
			Find(&dbRaces)
	})

}

func (r RacesRepository) GetRacesByClass(class string) ([]models.Race, error) {
	return r.selectRacesWithQuery(func(dbRaces *[]Race) *gorm.DB {
		return r.Db.
			Select("races.name, races.datetime", "championships.name AS championship_name", "championship.year AS championship_year", "tracks.name AS track", "layouts.name AS layout", "tracks.nation as track_nation", "tracks.location as track_location").
			Joins("join championship_classes ON races.championship = championship_classes.championship").
			Joins("join championships ON championships.id = races.championship").
			Joins("join layouts ON layouts.id = races.layout").
			Joins("join tracks ON layouts.track = tracks.name").
			Where("class = ?", class).
			Find(&dbRaces)
	})
}

func (r RacesRepository) GetRacesByTeam(teamName string) ([]models.Race, error) {
	return r.selectRacesWithQuery(func(dbRaces *[]Race) *gorm.DB {
		return r.Db.
			Distinct().
			Select("races.name, races.datetime", "championships.name AS championship_name", "championship.year AS championship_year", "tracks.name AS track", "layouts.name AS layout", "tracks.nation as track_nation", "tracks.location as track_location").
			Joins("join championships on races.championship = championships.id").
			Joins("join entries on championships.id = entries.championship").
			Joins("join layouts ON layouts.id = races.layout").
			Joins("join tracks ON layouts.track = tracks.name").
			Where("team = ?", teamName).
			Find(&dbRaces)
	})
}

func (r RacesRepository) GetDriversRacesByNationality(nation string) ([]models.Race, error) {
	return r.selectRacesWithQuery(func(dbRaces *[]Race) *gorm.DB {
		return r.Db.
			Distinct().
			Select("races.name, races.datetime", "championships.name AS championship_name", "championship.year AS championship_year", "tracks.name AS track", "layouts.name AS layout", "tracks.nation as track_nation", "tracks.location as track_location").
			Joins("join championships on races.championship = championships.id").
			Joins("join entries on championships.id = entries.championship").
			Joins("join driver_entries on  driver_entries.entry = entries.id").
			Joins("join drivers on drivers.cf = driver_entries.driver").
			Joins("join layouts ON layouts.id = races.layout").
			Joins("join tracks ON layouts.track = tracks.name").
			Where("drivers.nation = ?", nation).
			Find(&dbRaces)
	})
}

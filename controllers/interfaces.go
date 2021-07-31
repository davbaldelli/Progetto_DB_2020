package controllers

import "ProgettoDB/models"

type RacesController interface {
	GetIncomingRacesByClass(class string) ([]models.Race, error)
	GetIncomingRacesByTeam(teamName string) ([]models.Race, error)
	GetDriversRacesByNationality(nation string) ([]models.Race, error)
}

type ChampionshipsController interface {
	GetDriverChampionships(driver models.Driver) ([]models.Championship, error)
	GetIncomingChampionshipsByTeam(team models.Team) ([]models.Championship, error)
	GetDriversChampionshipsByNationality(nation string) ([]models.Championship, error)
}

type EntriesController interface {
	GetEntryByRaceNumber(championship models.Championship, raceNumber uint) (models.Entry, error)
}

type TeamController interface {
	GetTeamsWithoutParticipationByYear(year uint) ([]models.Team, error)
}

type StatisticsController interface {
	GetBrandCarsUsage(brandName string) ([]models.CarUsage, error)
	GetTrackLayoutsUsage(trackName string) ([]models.LayoutUsage, error)
}

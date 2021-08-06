package controllers

import "ProgettoDB/models"

type RacesController interface {
	GetIncomingRacesByClass(class string) ([]models.Race, error)
	GetIncomingRacesByTeam(teamName string) ([]models.Race, error)
	GetDriversRacesByNationality(nation string) ([]models.Race, error)
	GetChampionshipRaces(championship models.Championship) ([]models.Race, error)
}

type ChampionshipsController interface {
	GetDriverChampionships(driver models.Driver) ([]models.Championship, error)
	GetIncomingChampionshipsByTeam(team models.Team) ([]models.Championship, error)
	GetDriversChampionshipsByNationality(nation string) ([]models.Championship, error)
}

type EntriesController interface {
	GetEntryByRaceNumber(championship models.Championship, raceNumber uint) (models.Entry, error)
	GetChampionshipEntryList(championship models.Championship) ([]models.Entry, error)
}

type TeamController interface {
	GetAllTeams() ([]models.Team, error)
	GetTeamsWithoutParticipationByYear(year uint) ([]models.Team, error)
}

type StatisticsController interface {
	GetBrandCarsUsage(brandName string) ([]models.CarUsage, error)
	GetTrackLayoutsUsage(trackName string) ([]models.LayoutUsage, error)
}

type ClassesController interface {
	GetAllCLasses() ([]models.CarClass, error)
}

type DriversController interface {
	GetAllDrivers() ([]models.Driver, error)
}

type NationsController interface {
	GetAllNations() ([]string, error)
}

type ManufacturersController interface {
	GetAllManufacturers() ([]models.Brand, error)
}

type TracksController interface {
	GetAllTracks() ([]models.Track, error)
}

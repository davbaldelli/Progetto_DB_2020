package controllers

import "ProgettoDB/models"

type RacesController interface {
	GetRacesByClass(class string) ([]models.Race, error)
	GetRacesByTeam(teamName string) ([]models.Race, error)
	GetDriversRacesByNationality(nation string) ([]models.Race, error)
	GetChampionshipRaces(championship models.Championship) ([]models.Race, error)
	GetRaceResult(race models.Race) ([]models.Result, error)
	GetAllRaces() ([]models.Race, error)
}

type ChampionshipsController interface {
	GetDriverChampionships(driver models.Driver) ([]models.Championship, error)
	GetChampionshipsByTeam(team models.Team) ([]models.Championship, error)
	GetDriversChampionshipsByNationality(nation string) ([]models.Championship, error)
	GetAllChampionships() ([]models.Championship, error)
	GetChampionshipsByClass(class string) ([]models.Championship, error)
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
	GetTheFiveMostUsedTracks() ([]models.TrackUsage, error)
}

type ClassesController interface {
	GetAllClasses() ([]models.CarClass, error)
}

type DriversController interface {
	GetAllDrivers() ([]models.Driver, error)
	GetFiveDriversWithMoreRaces() ([]models.DriverRaces, error)
	InsertDriver(driver models.Driver) error
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

type CarController interface {
	GetChampionshipCars(models.Championship) ([]models.Car, error)
	GetDriverCarOnCircuit(models.Driver, models.Track) ([]models.Car, error)
	GetAllCars() ([]models.Car, error)
}

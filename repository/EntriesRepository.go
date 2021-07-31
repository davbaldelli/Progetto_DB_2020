package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
	"time"
)

type EntriesRepository struct {
	Db *gorm.DB
}

type DriverEntry struct {
	Id           uint
	Championship uint
	RaceNumber   uint
	Team         string
	Model        string
	Year         uint
	Brand        string
	Class        string
	Drivetrain   models.Drivetrain
	Transmission models.Transmission
	Cf           string
	Name         string
	Surname      string
	Birthdate    time.Time
	Nation       string
	Sex          models.Sex
}

func (e EntriesRepository) GetEntryByRaceNumber(championship models.Championship, raceNumber uint) (models.Entry, error) {
	var dbEntries []DriverEntry
	if err := e.Db.Table("championships").
		Select("entries.*", "cars.*", "drivers.*").
		Where("championships.name = ? AND championships.year = ?", championship.Name, championship.Year).
		Joins("join entries ON entries.championship = championships.id").
		Joins("JOIN cars ON cars.id = entries.car").
		Joins("JOIN driver_entries ON driver_entries.entry = entries.id").
		Joins("JOIN drivers ON driver_entries.driver = drivers.cf").
		Where("entries.race_number = ?", raceNumber).
		Find(&dbEntries).Error; err != nil {
		return models.Entry{}, err
	}
	var entry = models.Entry{
		Car: models.Car{
			Model:        dbEntries[0].Model,
			Year:         dbEntries[0].Year,
			Brand:        models.Brand{Name: dbEntries[0].Brand},
			Class:        dbEntries[0].Class,
			Drivetrain:   dbEntries[0].Drivetrain,
			Transmission: dbEntries[0].Transmission,
		},
		RaceNumber: dbEntries[0].RaceNumber,
		Team:       models.Team{Name: dbEntries[0].Team},
	}
	for _, dbEntry := range dbEntries {
		entry.Drivers = append(entry.Drivers, models.Driver{
			Name:      dbEntry.Name,
			Surname:   dbEntry.Surname,
			CF:        dbEntry.Cf,
			Sex:       dbEntry.Sex,
			Birthdate: dbEntry.Birthdate,
		})
	}
	return entry, nil
}

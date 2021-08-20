package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type ChampionshipRepository struct {
	Db *gorm.DB
}

type championshipQuery func(*[]Championship) *gorm.DB

func dbChampionshipToEntity(championship Championship) models.Championship {
	var classes []models.CarClass
	for _, class := range championship.Classes {
		classes = append(classes, models.CarClass{Name: class.Name})
	}
	return models.Championship{
		Name:      championship.Name,
		Year:      championship.Year,
		EntryList: nil,
		Races:     nil,
		Classes:   classes,
	}
}

func (c ChampionshipRepository) selectChampionshipsWithQuery(query championshipQuery) ([]models.Championship, error) {
	var dbChamps []Championship
	if err := query(&dbChamps).Error; err != nil {
		return nil, err
	}

	var champs []models.Championship

	for _, dbChamp := range dbChamps {
		var dbClasses []CarClass
		if err := c.Db.Table("championship_classes").
			Select("championship_classes.class AS name").
			Where("championship_classes.championship = ?", dbChamp.Id).
			Find(&dbClasses).Error; err != nil {
			return nil, err
		}
		dbChamp.Classes = dbClasses
		champs = append(champs, dbChampionshipToEntity(dbChamp))
	}

	return champs, nil
}

func (c ChampionshipRepository) GetAllChampionships() ([]models.Championship, error) {
	return c.selectChampionshipsWithQuery(func(dbChamps *[]Championship) *gorm.DB {
		return c.Db.Find(&dbChamps)
	})
}

func (c ChampionshipRepository) GetDriverChampionships(driver models.Driver) ([]models.Championship, error) {
	return c.selectChampionshipsWithQuery(func(dbChamps *[]Championship) *gorm.DB {
		return c.Db.Distinct().
			Select("championships.*").
			Joins("join entries on championships.id = entries.championship").
			Joins("join driver_entries on  driver_entries.entry = entries.id").
			Joins("join drivers on drivers.cf = driver_entries.driver").
			Where("drivers.cf = ?", driver.CF).
			Find(&dbChamps)
	})
}

func (c ChampionshipRepository) GetChampionshipsByTeam(team models.Team) ([]models.Championship, error) {
	return c.selectChampionshipsWithQuery(func(dbChamps *[]Championship) *gorm.DB {
		return c.Db.Distinct().
			Where("entries.team = ?", team.Name).
			Select("championships.*").
			Joins("join entries on championships.id = entries.championship").
			Find(&dbChamps)
	})
}

func (c ChampionshipRepository) GetDriversChampionshipsByNationality(nation string) ([]models.Championship, error) {
	return c.selectChampionshipsWithQuery(func(dbChamps *[]Championship) *gorm.DB {
		return c.Db.Distinct().
			Joins("join entries on championships.id = entries.championship").
			Joins("join driver_entries on  driver_entries.entry = entries.id").
			Joins("join drivers on drivers.cf = driver_entries.driver").
			Where("drivers.nation = ?", nation).
			Find(&dbChamps)
	})
}

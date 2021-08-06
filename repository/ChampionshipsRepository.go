package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type ChampionshipRepository struct {
	Db *gorm.DB
}

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

func (c ChampionshipRepository) GetDriverChampionships(driver models.Driver) ([]models.Championship, error) {
	var dbChamps []Championship
	if err := c.Db.Table("drivers").Distinct().
		Where("drivers.cf = ?", driver.CF).
		Select("championships.*").
		Joins("join driver_entries on drivers.cf = driver_entries.driver").
		Joins("join entries on  driver_entries.entry = entries.id").
		Joins("join championships on championships.id = entries.championship").
		Preload("Races").
		Preload("Entries").
		Find(&dbChamps).Error; err != nil {
		return nil, err
	}

	var champs []models.Championship

	for _, dbChamp := range dbChamps {
		champs = append(champs, dbChampionshipToEntity(dbChamp))
	}
	return champs, nil
}

func (c ChampionshipRepository) GetIncomingChampionshipsByTeam(team models.Team) ([]models.Championship, error) {
	var dbChamps []Championship
	if err := c.Db.Table("entries").Distinct().
		Where("entries.team = ?", team.Name).
		Select("championships.*").
		Joins("join championships on championships.id = entries.championship").
		Preload("Races").
		Preload("Entries").
		Find(&dbChamps).Error; err != nil {
		return nil, err
	}

	var champs []models.Championship

	for _, dbChamp := range dbChamps {
		champs = append(champs, dbChampionshipToEntity(dbChamp))
	}
	return champs, nil
}

func (c ChampionshipRepository) GetDriversChampionshipsByNationality(nation string) ([]models.Championship, error) {
	var dbChamps []Championship
	if err := c.Db.Table("drivers").Distinct().
		Where("drivers.nation = ?", nation).
		Select("championships.*").
		Joins("join driver_entries on drivers.cf = driver_entries.driver").
		Joins("join entries on  driver_entries.entry = entries.id").
		Joins("join championships on championships.id = entries.championship").
		Preload("Races").
		Preload("Entries").
		Preload("Classes").
		Find(&dbChamps).Error; err != nil {
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

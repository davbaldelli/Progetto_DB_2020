package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type ChampionshipRepository struct {
	Db *gorm.DB
}

func dbChampionshipToEntity(championship Championship) models.Championship {
 return models.Championship{
	 Name:      championship.Name,
	 Year:      championship.Year,
	 EntryList: nil,
	 Races:     nil,
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
		Find(&dbChamps).Error ; err != nil{
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
		Find(&dbChamps).Error ; err != nil{
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
		Find(&dbChamps).Error ; err != nil{
		return nil, err
	}

	var champs []models.Championship

	for _, dbChamp := range dbChamps {
		champs = append(champs, dbChampionshipToEntity(dbChamp))
	}
	return champs, nil
}



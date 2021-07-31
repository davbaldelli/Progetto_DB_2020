package main

import (
	"ProgettoDB/models"
	"ProgettoDB/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "davide:1908FCInter@tcp(192.168.122.76:3306)/progetto_db?charset=utf8mb4&parseTime=True&loc=Local"
	dbase, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err1 != nil {
		log.Fatal(err1)
	}
	racesRepository := repository.RacesRepository{Db: dbase}
	championshipsRepo := repository.ChampionshipRepository{Db: dbase}
	entriesRepo := repository.EntriesRepository{Db: dbase}

	if races, err := racesRepository.GetIncomingRacesByTeam("Frikadelli Racing Team"); err != nil {
		log.Fatal(err)
	} else {
		log.Print(races)
	}

	if champs, err := championshipsRepo.GetDriversChampionshipsByNationality("Italy"); err != nil {
		log.Fatal(err)
	} else {
		log.Print(champs)
	}

	if entry, err := entriesRepo.GetEntryByRaceNumber(models.Championship{
		Name: "GT World Challenge Europe",
		Year: 2021,
	}, 51); err != nil {
		log.Fatal(err)
	} else {
		log.Print(entry)
	}

}

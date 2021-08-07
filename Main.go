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
	statisticsRepo := repository.StatisticsRepository{Db: dbase}
	classesRepo := repository.ClassesRepository{Db: dbase}
	driversRepo := repository.DriverRepository{Db: dbase}
	brandsRepo := repository.ManufacturerRepository{Db: dbase}
	tracksRepo := repository.TracksRepository{Db: dbase}

	if races, err := racesRepository.GetChampionshipRaces(models.Championship{Name: "GT World Challenge Europe", Year: 2021}); err != nil {
		log.Fatal(err)
	} else {
		log.Print(races)
	}

	if champs, err := championshipsRepo.GetDriversChampionshipsByNationality("Italy"); err != nil {
		log.Fatal(err)
	} else {
		log.Print(champs)
	}

	if entry, err := entriesRepo.GetChampionshipEntryList(models.Championship{Name: "GT World Challenge Europe", Year: 2021}); err != nil {
		log.Fatal(err)
	} else {
		log.Print(entry)
	}

	if stats, err := statisticsRepo.GetBrandCarsUsage("Ferrari"); err != nil {
		log.Fatal(err)
	} else {
		log.Print(stats)
	}

	if classes, err := classesRepo.GetAllCLasses(); err != nil {
		log.Fatal(err)
	} else {
		log.Print(classes)
	}

	if drivers, err := driversRepo.GetAllDrivers(); err != nil {
		log.Fatal(err)
	} else {
		log.Print(drivers)
	}

	if brands, err := brandsRepo.GetAllManufacturers(); err != nil {
		log.Fatal(err)
	} else {
		log.Print(brands)
	}

	if tracks, err := tracksRepo.GetAllTracks(); err != nil {
		log.Fatal(err)
	} else {
		log.Print(tracks)
	}

}

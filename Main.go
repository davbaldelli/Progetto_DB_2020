package main

import (
	"ProgettoDB/repository"
	"ProgettoDB/routes"
	"ProgettoDB/routes/handlers"
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
	racesHandler := handlers.RacesHandler{Ctrl: repository.RacesRepository{Db: dbase}}
	championshipsHandler := handlers.ChampionshipsHandler{Ctrl: repository.ChampionshipRepository{Db: dbase}}
	//entriesRepo := repository.EntriesRepository{Db: dbase}
	//statisticsRepo := repository.StatisticsRepository{Db: dbase}
	//classesRepo := repository.ClassesRepository{Db: dbase}
	//driversRepo := repository.DriverRepository{Db: dbase}
	//brandsRepo := repository.ManufacturerRepository{Db: dbase}
	//tracksRepo := repository.TracksRepository{Db: dbase}
	//teamsRepo := repository.TeamsRepository{Db: dbase}
	//carsRepo := repository.CarRepository{Db: dbase}

	web := routes.Web{ChampionshipsHandler: championshipsHandler, RacesHandler: racesHandler}

	web.Listen()

}

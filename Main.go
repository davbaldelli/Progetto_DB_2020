package main

import (
	"ProgettoDB/repository"
	"ProgettoDB/routes"
	"ProgettoDB/routes/handlers"
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

type Credentials struct {
	Username string
	Password string
	DBIP     string
}

func main() {

	jsonFile, err := os.Open("credentials.json")

	if err != nil {
		log.Fatal(err)
	}

	byteValue, err1 := ioutil.ReadAll(jsonFile)

	if err1 != nil {
		log.Fatal(err)
	}

	var cred Credentials

	if err2 := json.Unmarshal(byteValue, &cred); err2 != nil {
		log.Fatal(err)
	}

	dsn := cred.Username + ":" + cred.Password + "@tcp(" + cred.DBIP + ":3306)/progetto_db?charset=utf8mb4&parseTime=True&loc=Local"
	dbase, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err1 != nil {
		log.Fatal(err1)
	} else {
		log.Print("Connected to database successfully")
	}

	racesHandler := handlers.RacesHandler{Ctrl: repository.RacesRepository{Db: dbase}}
	championshipsHandler := handlers.ChampionshipsHandler{Ctrl: repository.ChampionshipRepository{Db: dbase}}
	entriesHandler := handlers.EntriesHandler{Ctrl: repository.EntriesRepository{Db: dbase}}
	statisticsHandler := handlers.StatisticsHandler{Ctrl: repository.StatisticsRepository{Db: dbase}}
	classesHandler := handlers.ClassesHandler{Ctrl: repository.ClassesRepository{Db: dbase}}
	driversHandler := handlers.DriversHandler{Ctrl: repository.DriverRepository{Db: dbase}}
	brandsHandler := handlers.ManufacturersHandler{Ctrl: repository.ManufacturerRepository{Db: dbase}}
	teamsHandler := handlers.TeamsHandler{Ctrl: repository.TeamsRepository{Db: dbase}}
	carsHandler := handlers.CarHandler{Ctrl: repository.CarRepository{Db: dbase}}
	tracksHandler := handlers.TracksHandler{Ctrl: repository.TracksRepository{Db: dbase}}
	nationsHandler := handlers.NationsHandler{Ctrl: repository.NationsRepository{Db: dbase}}

	web := routes.Web{
		ChampionshipsHandler: championshipsHandler,
		RacesHandler:         racesHandler,
		EntriesHandler:       entriesHandler,
		TeamsHandler:         teamsHandler,
		StatisticsHandler:    statisticsHandler,
		CarHandler:           carsHandler,
		ClassesHandler:       classesHandler,
		DriversHandler:       driversHandler,
		ManufacturersHandler: brandsHandler,
		TracksHandler:        tracksHandler,
		NationsHandler:       nationsHandler,
	}

	web.Listen()

}

package routes

import (
	"ProgettoDB/routes/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Router interface {
	Listen()
}

type Web struct {
	ChampionshipsHandler handlers.ChampionshipsHandler
	RacesHandler         handlers.RacesHandler
	EntriesHandler       handlers.EntriesHandler
	TeamsHandler         handlers.TeamsHandler
	StatisticsHandler    handlers.StatisticsHandler
	CarHandler           handlers.CarHandler
	ClassesHandler       handlers.ClassesHandler
	TracksHandler        handlers.TracksHandler
	ManufacturersHandler handlers.ManufacturersHandler
	DriversHandler       handlers.DriversHandler
}

func (w Web) Listen() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/championship/all", w.ChampionshipsHandler.GETAllChampionships).Methods("GET")
	router.HandleFunc("/championship/driver/{cf}", w.ChampionshipsHandler.GETDriverChampionships).Methods("GET")
	router.HandleFunc("/championship/team/{team}", w.ChampionshipsHandler.GETTeamChampionships).Methods("GET")
	router.HandleFunc("/championship/nation/driver/{nation}", w.ChampionshipsHandler.GETDriversChampionshipsByNation).Methods("GET")
	router.HandleFunc("/championship/class/{class}", w.ChampionshipsHandler.GETChampionshipsByClass).Methods("GET")

	router.HandleFunc("/race/championship/{name}/{year}", w.RacesHandler.GETRacesByChampionship).Methods("GET")
	router.HandleFunc("/race/team/{team}", w.RacesHandler.GETRacesByTeam).Methods("GET")
	router.HandleFunc("/race/nation/driver/{nation}", w.RacesHandler.GETRacesByDriversNation).Methods("GET")
	router.HandleFunc("/race/class/{class}", w.RacesHandler.GETRacesByClass).Methods("GET")

	router.HandleFunc("/entry/championship/{name}/{year}", w.EntriesHandler.GETChampionshipEntries).Methods("GET")
	router.HandleFunc("/entry/championship/{name}/{year}/{number}", w.EntriesHandler.GETEntryByRaceNumber).Methods("GET")

	router.HandleFunc("/team/all", w.TeamsHandler.GETAllTeams).Methods("GET")
	router.HandleFunc("/team/noparticipation/{year}", w.TeamsHandler.GETTeamsWithoutParticipationByYear).Methods("GET")

	router.HandleFunc("/statistic/track/layout/usage/{track}", w.StatisticsHandler.GETTrackLayoutsUsage).Methods("GET")
	router.HandleFunc("/statistic/brand/car/usage/{brand}", w.StatisticsHandler.GETBrandsCarsUsage).Methods("GET")
	router.HandleFunc("/statistics/mostusedtracks", w.StatisticsHandler.GETMostFiveUsedTracks).Methods("GET")

	router.HandleFunc("/car/championship/{name}/{year}", w.CarHandler.GETChampionshipCars).Methods("GET")
	router.HandleFunc("/car/driver/track/{driver}/{track}", w.CarHandler.GETDriverCarsOnTrack).Methods("GET")

	router.HandleFunc("/class/all", w.ClassesHandler.GETAllClasses).Methods("GET")

	router.HandleFunc("/drivers/all", w.DriversHandler.GETAllDrivers).Methods("GET")
	router.HandleFunc("/drivers/fivewithmoreraces", w.DriversHandler.GETFiveDriversWithMoreRaces).Methods("GET")

	router.HandleFunc("/tracks/all", w.TracksHandler.GETAllTracks).Methods("GET")

	router.HandleFunc("/manufacturer/all", w.ManufacturersHandler.GETAllManufacturers).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("127.0.0.1"+
		":1234", handler))

}

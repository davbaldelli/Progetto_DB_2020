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
}

func (w Web) Listen() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/championships/all", w.ChampionshipsHandler.GETAllChampionships).Methods("GET")
	router.HandleFunc("/championships/driver/{cf}", w.ChampionshipsHandler.GETDriverChampionships).Methods("GET")
	router.HandleFunc("/championships/team/{team}", w.ChampionshipsHandler.GETTeamChampionships).Methods("GET")
	router.HandleFunc("/championships/nation/driver/{nation}", w.ChampionshipsHandler.GETDriversChampionshipsByNation).Methods("GET")

	router.HandleFunc("/race/championship/{name}/{year}", w.RacesHandler.GETRacesByChampionship).Methods("GET")
	router.HandleFunc("/race/team/{team}", w.RacesHandler.GETRacesByTeam).Methods("GET")
	router.HandleFunc("/race/nation/driver/{nation}", w.RacesHandler.GETRacesByDriversNation).Methods("GET")
	router.HandleFunc("/race/class/{class}", w.RacesHandler.GETRacesByClass).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("127.0.0.1"+
		":1234", handler))

}

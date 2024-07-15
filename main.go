package main

import (
	"log"
	"net/http"

	"Futbol_Sim/class"
	"Futbol_Sim/handlers"
	"Futbol_Sim/model"
	"Futbol_Sim/service"
)

func main() {
	model.InitDB()

	teamNames := []string{"Liverpool", "Manchester United", "Manchester City", "Chelsea"}
	league := class.NewLeague(teamNames)

	leagueService := service.NewLeagueService(league)
	leagueHandler := handlers.NewLeagueHandler(leagueService)

	teamService := service.NewTeamService()
	teamHandler := handlers.NewTeamHandler(teamService)

	matchService := service.NewMatchService()
	matchHandler := handlers.NewMatchHandler(matchService)

	http.HandleFunc("/simulate", leagueHandler.SimulateHandler)
	http.HandleFunc("/next-week", leagueHandler.NextWeekHandler)
	http.HandleFunc("/standings", leagueHandler.StandingsHandler)
	http.HandleFunc("/matches", matchHandler.GetMatchesHandler)
	http.HandleFunc("/teams", teamHandler.GetTeamsHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

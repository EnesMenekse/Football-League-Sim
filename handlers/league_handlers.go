package handlers

import (
	"Futbol_Sim/service"
	"encoding/json"
	"net/http"
)

// LeagueHandler, lig işlemlerini yürüten handler.
type LeagueHandler struct {
	service service.LeagueService
}

// NewLeagueHandler, yeni bir LeagueHandler örneği döndürür.
func NewLeagueHandler(service service.LeagueService) *LeagueHandler {
	return &LeagueHandler{service: service}
}

func (lh *LeagueHandler) SimulateHandler(w http.ResponseWriter, r *http.Request) {
	err := lh.service.Simulate()
	if err != nil {
		http.Error(w, "Unable to simulate league", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"league_ended": true})
}

func (lh *LeagueHandler) NextWeekHandler(w http.ResponseWriter, r *http.Request) {
	err := lh.service.NextWeek()
	if err != nil {
		http.Error(w, "Unable to process next week", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"league_ended": false})
}

func (lh *LeagueHandler) StandingsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := lh.service.GetStandings()
	if err != nil {
		http.Error(w, "Unable to get standings", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func (lh *LeagueHandler) MatchesHandler(w http.ResponseWriter, r *http.Request) {
	matches, err := lh.service.GetMatches()
	if err != nil {
		http.Error(w, "Unable to get matches", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

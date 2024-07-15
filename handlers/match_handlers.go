package handlers

import (
	"Futbol_Sim/service"
	"encoding/json"
	"net/http"
)

// MatchHandler, maçlarla ilgili işlemleri yürüten handler.
type MatchHandler struct {
	service service.MatchService
}

// NewMatchHandler, yeni bir MatchHandler örneği döndürür.
func NewMatchHandler(service service.MatchService) *MatchHandler {
	return &MatchHandler{service: service}
}

func (mh *MatchHandler) GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	matches, err := mh.service.GetMatches()
	if err != nil {
		http.Error(w, "Unable to get matches", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

package handlers

import (
	"Futbol_Sim/service"
	"encoding/json"
	"net/http"
)

// TeamHandler, takımlarla ilgili işlemleri yürüten handler.
type TeamHandler struct {
	service service.TeamService
}

// NewTeamHandler, yeni bir TeamHandler örneği döndürür.
func NewTeamHandler(service service.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

func (th *TeamHandler) GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := th.service.GetTeams()
	if err != nil {
		http.Error(w, "Unable to get teams", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

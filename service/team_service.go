package service

import (
	"Futbol_Sim/class"
	"Futbol_Sim/model"
)

// TeamService arayüzü, takım işlemlerini tanımlar.
type TeamService interface {
	GetTeams() ([]*class.Team, error)
}

// teamService, TeamService arayüzünü uygular.
type teamService struct{}

// NewTeamService, yeni bir TeamService örneği döndürür.
func NewTeamService() TeamService {
	return &teamService{}
}

func (ts *teamService) GetTeams() ([]*class.Team, error) {
	return model.GetTeams()
}

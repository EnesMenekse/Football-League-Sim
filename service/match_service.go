package service

import (
	"Futbol_Sim/class"
	"Futbol_Sim/model"
)

// MatchService arayüzü, maç işlemlerini tanımlar.
type MatchService interface {
	GetMatches() ([]*class.Match, error)
}

// matchService, MatchService arayüzünü uygular.
type matchService struct{}

// NewMatchService, yeni bir MatchService örneği döndürür.
func NewMatchService() MatchService {
	return &matchService{}
}

func (ms *matchService) GetMatches() ([]*class.Match, error) {
	matches, err := model.GetMatches()
	if err != nil {
		return nil, err
	}
	var result []*class.Match
	for _, m := range matches {
		result = append(result, &class.Match{
			HomeTeam:  &class.Team{Name: m.HomeTeam},
			AwayTeam:  &class.Team{Name: m.AwayTeam},
			HomeGoals: m.HomeGoals,
			AwayGoals: m.AwayGoals,
		})
	}
	return result, nil
}

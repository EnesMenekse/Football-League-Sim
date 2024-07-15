package service

import (
	"Futbol_Sim/class"
	"Futbol_Sim/model"
)

// LeagueService arayüzü, lig işlemlerini tanımlar.
type LeagueService interface {
	Simulate() error
	NextWeek() error
	GetStandings() ([]*class.Team, error)
	GetMatches() ([]*class.Match, error)
}

// leagueService, LeagueService arayüzünü uygular.
type leagueService struct {
	league *class.League
}

// NewLeagueService, yeni bir LeagueService örneği döndürür.
func NewLeagueService(league *class.League) LeagueService {
	return &leagueService{league: league}
}

func (ls *leagueService) Simulate() error {
	for ls.league.CurrentWeek < len(ls.league.Weeks) {
		ls.league.PlayNextWeek(model.SaveTeamStats)
	}
	model.SaveMatches(ls.league.Matches)
	return nil
}

func (ls *leagueService) NextWeek() error {
	ls.league.PlayNextWeek(model.SaveTeamStats)
	return nil
}

func (ls *leagueService) GetStandings() ([]*class.Team, error) {
	return model.GetTeams()
}

func (ls *leagueService) GetMatches() ([]*class.Match, error) {
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

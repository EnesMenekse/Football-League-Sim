package model

import (
	"Futbol_Sim/class"
	"log"
)

func SaveMatch(match *class.Match) {
	_, err := DB.Exec(`
        INSERT INTO matches (home_team_id, away_team_id, home_goals, away_goals, week)
        VALUES ((SELECT id FROM teams WHERE name = $1), (SELECT id FROM teams WHERE name = $2), $3, $4, $5)`,
		match.HomeTeam.Name, match.AwayTeam.Name, match.HomeGoals, match.AwayGoals, match.Week)
	if err != nil {
		log.Println("Error saving match:", err)
	}
}

func SaveMatches(matches []*class.Match) {
	for _, match := range matches {
		SaveMatch(match)
	}
}

func GetMatches() ([]struct {
	HomeTeam  string
	AwayTeam  string
	HomeGoals int
	AwayGoals int
}, error) {
	rows, err := DB.Query(`
        SELECT
            ht.name as home_team,
            at.name as away_team,
            m.home_goals,
            m.away_goals
        FROM matches m
        JOIN teams ht ON m.home_team_id = ht.id
        JOIN teams at ON m.away_team_id = at.id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []struct {
		HomeTeam  string
		AwayTeam  string
		HomeGoals int
		AwayGoals int
	}
	for rows.Next() {
		var score struct {
			HomeTeam  string
			AwayTeam  string
			HomeGoals int
			AwayGoals int
		}
		if err := rows.Scan(&score.HomeTeam, &score.AwayTeam, &score.HomeGoals, &score.AwayGoals); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scores, nil
}

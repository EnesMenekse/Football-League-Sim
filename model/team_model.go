package model

import (
	"Futbol_Sim/class"
	"log"
)

func SaveTeamStats(team *class.Team) {
	_, err := DB.Exec(`
        UPDATE teams
        SET points = $1, wins = $2, draws = $3, losses = $4, goals_for = $5, goals_against = $6
        WHERE name = $7`,
		team.Points, team.Wins, team.Draws, team.Losses, team.GoalsFor, team.GoalsAgainst, team.Name)
	if err != nil {
		log.Println("Error updating team stats:", err)
	}
}

func InsertTeam(team *class.Team) {
	if !TeamExists(team.Name) {
		_, err := DB.Exec(`
            INSERT INTO teams (name, points, wins, draws, losses, goals_for, goals_against)
            VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			team.Name, team.Points, team.Wins, team.Draws, team.Losses, team.GoalsFor, team.GoalsAgainst)
		if err != nil {
			log.Println("Error inserting team:", err)
		}
	}
}

func TeamExists(name string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM teams WHERE name=$1)`
	err := DB.QueryRow(query, name).Scan(&exists)
	if err != nil {
		log.Println("Error checking if team exists:", err)
	}
	return exists
}

func GetTeams() ([]*class.Team, error) {
	rows, err := DB.Query(`SELECT id, name, points, wins, draws, losses, goals_for, goals_against FROM teams`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []*class.Team
	for rows.Next() {
		var team class.Team
		if err := rows.Scan(&team.ID, &team.Name, &team.Points, &team.Wins, &team.Draws, &team.Losses, &team.GoalsFor, &team.GoalsAgainst); err != nil {
			return nil, err
		}
		teams = append(teams, &team)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

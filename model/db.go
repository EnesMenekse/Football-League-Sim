package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "host=db user=postgres password=1234 dbname=FootballLeague sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	clearTables()
}

func clearTables() {
	_, err := DB.Exec("DELETE FROM matches")
	if err != nil {
		log.Println("Error clearing matches table:", err)
	}

	_, err = DB.Exec("DELETE FROM teams")
	if err != nil {
		log.Println("Error clearing teams table:", err)
	}

	_, err = DB.Exec("ALTER SEQUENCE teams_id_seq RESTART WITH 1")
	if err != nil {
		log.Println("Error resetting team ID sequence:", err)
	}

	_, err = DB.Exec("ALTER SEQUENCE matches_id_seq RESTART WITH 1")
	if err != nil {
		log.Println("Error resetting match ID sequence:", err)
	}
}

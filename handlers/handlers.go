package handlers

import (
	"Futbol_Sim/class"
)

var league *class.League

func SetLeague(l *class.League) {
	league = l
}

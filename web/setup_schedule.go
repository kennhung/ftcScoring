package web

import (
	"net/http"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"github.com/kennhung/ftcScoring/tournament"
	"github.com/kennhung/ftcScoring/model"
)

func (web *Web) setupscheduleGETHandler(w http.ResponseWriter, r *http.Request) {

	var currentMatch = new(model.Match)
	var allMatchs [3][]model.Match
	var err error

	allMatchs[0], err = web.arena.Database.GetMatchesByType("practice")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	allMatchs[1], err = web.arena.Database.GetMatchesByType("qualification")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	allMatchs[2], err = web.arena.Database.GetMatchesByType("elimination")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	buffer := new(bytes.Buffer)
	template.Setup_Schedule(allMatchs,currentMatch,buffer)
	w.Write(buffer.Bytes())
}

func (web *Web) setupschedulePOSTHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := web.arena.Database.GetAllTeams()
	if err != nil {
		handleWebErr(w, err)
		return
	}

	matches, err := tournament.BuildRandomSchedule(teams, 3, "practice")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	for _, match := range matches {
		err = web.arena.Database.CreateMatch(&match)
		if err != nil {
			handleWebErr(w, err)
			return
		}
	}
}
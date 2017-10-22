package web

import (
	"net/http"
	"github.com/kennhung/ftcScoring/scheduling"
)

func (web *Web) setupscheduleGETHandler(w http.ResponseWriter, r *http.Request) {

	teams, err := web.arena.Database.GetAllTeams()
	if err != nil {
		handleWebErr(w, err)
		return
	}

	matchs, err := scheduling.BuildRandomSchedule(teams, 3, "")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	for _, match := range matchs {
		web.arena.Database.CreateMatch(&match)
	}

}

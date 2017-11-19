package web

import (
	"net/http"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"github.com/kennhung/ftcScoring/scheduling"
)

func (web *Web) setupscheduleGETHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	template.Setup_Schedule(buffer)
	w.Write(buffer.Bytes())
}

func (web *Web) setupschedulePOSTHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := web.arena.Database.GetAllTeams()
	if err != nil {
		handleWebErr(w, err)
		return
	}

	matches, err := scheduling.BuildRandomSchedule(teams, 3, "practice")
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
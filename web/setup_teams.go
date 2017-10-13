package web

import (
	"net/http"
	"github.com/kennhung/ftcScoring/webTemplate"
	"bytes"
)

// Shows the team list.
func (web *Web) teamsGetHandler(w http.ResponseWriter, r *http.Request) {

	web.renderTeams(w, r, false)
}

func (web *Web) renderTeams(w http.ResponseWriter, r *http.Request, showErrorMessage bool) {
	teams, err := web.arena.Database.GetAllTeams()
	if err != nil {
		handleWebErr(w, err)
		return
	}

	buffer := new(bytes.Buffer)
	template.Setup_Teams(teams,buffer)

	w.Write(buffer.Bytes())
}
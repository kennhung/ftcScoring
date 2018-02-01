package web

import (
	"net/http"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"time"
)

func (web *Web) setupsettingsGetHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)

	template.Setup_settings(web.arena.EventSettings, buffer)

	w.Write(buffer.Bytes())
}

func (web *Web) setupsettingsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	eventSettings := web.arena.EventSettings

	eventSettings.Name = r.PostFormValue("name")
	eventSettings.Type = r.PostFormValue("type")
	eventSettings.Region = r.PostFormValue("region")
	datestr := r.PostFormValue("date")
	eventSettings.DisplayBackgroundColor = r.PostFormValue("background_color")

	t, err := time.Parse("01/02/2006", datestr)
	if err != nil {
		handleWebErr(w, err)
		return
	}
	eventSettings.Date = t

	err = web.arena.Database.SaveEventSettings(eventSettings)
	if err != nil {
		handleWebErr(w, err)
		return
	}

	err = web.arena.LoadSettings()
	if err != nil {
		handleWebErr(w, err)
		return
	}

	http.Redirect(w, r, "/setup/settings", 303)
}

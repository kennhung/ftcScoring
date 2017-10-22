package web

import (
	"github.com/Team254/cheesy-arena/game"
	"io"
	"log"
	"net/http"
	"github.com/kennhung/ftcScoring/model"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
)

// Renders the audience display to be chroma keyed over the video feed.
func (web *Web) audienceDisplayHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	if web.arena.EventSettings.DisplayOverlayMode {
		template.Display_Audience_overlay(web.arena.EventSettings,buffer)
	}else{
		template.Display_Audience_NoOverlay(web.arena.EventSettings,buffer)
	}

	w.Write(buffer.Bytes())
}

// The websocket endpoint for the audience display client to receive status updates.
func (web *Web) audienceDisplayWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	websocket, err := NewWebsocket(w, r)
	if err != nil {
		handleWebErr(w, err)
		return
	}
	defer websocket.Close()

	audienceDisplayListener := web.arena.AudienceDisplayChannel.Listen()
	defer close(audienceDisplayListener)
	matchLoadTeamsListener := web.arena.MatchLoadTeamsChannel.Listen()
	defer close(matchLoadTeamsListener)
	matchTimeListener := web.arena.MatchTimeChannel.Listen()
	defer close(matchTimeListener)
	scorePostedListener := web.arena.ScorePostedChannel.Listen()
	defer close(scorePostedListener)
	playSoundListener := web.arena.PlaySoundChannel.Listen()
	defer close(playSoundListener)
	allianceSelectionListener := web.arena.AllianceSelectionChannel.Listen()
	defer close(allianceSelectionListener)
	reloadDisplaysListener := web.arena.ReloadDisplaysChannel.Listen()
	defer close(reloadDisplaysListener)

	// Send the various notifications immediately upon connection.
	var data interface{}
	err = websocket.Write("matchTiming", game.MatchTiming)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}
	err = websocket.Write("matchTime", MatchTimeMessage{web.arena.MatchState, int(web.arena.MatchTimeSec())})
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}
	err = websocket.Write("setAudienceDisplay", web.arena.AudienceDisplayScreen)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}
	data = struct {
		Match     *model.Match
		MatchName string
	}{web.arena.CurrentMatch, web.arena.CurrentMatch.CapitalizedType()}
	err = websocket.Write("setMatch", data)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}

	// Spin off a goroutine to listen for notifications and pass them on through the websocket.
	go func() {
		for {
			var messageType string
			var message interface{}
			select {
			case _, ok := <-audienceDisplayListener:
				if !ok {
					return
				}
				messageType = "setAudienceDisplay"
				message = web.arena.AudienceDisplayScreen
			case _, ok := <-matchLoadTeamsListener:
				if !ok {
					return
				}
				messageType = "setMatch"
				message = struct {
					Match     *model.Match
					MatchName string
				}{web.arena.CurrentMatch, web.arena.CurrentMatch.CapitalizedType()}
			case matchTimeSec, ok := <-matchTimeListener:
				if !ok {
					return
				}
				messageType = "matchTime"
				message = MatchTimeMessage{web.arena.MatchState, matchTimeSec.(int)}
			case sound, ok := <-playSoundListener:
				if !ok {
					return
				}
				messageType = "playSound"
				message = sound
			case _, ok := <-reloadDisplaysListener:
				if !ok {
					return
				}
				messageType = "reload"
				message = nil
			}
			err = websocket.Write(messageType, message)
			if err != nil {
				// The client has probably closed the connection; nothing to do here.
				return
			}
		}
	}()

	// Loop, waiting for commands and responding to them, until the client closes the connection.
	for {
		_, _, err := websocket.Read()
		if err != nil {
			if err == io.EOF {
				// Client has closed the connection; nothing to do here.
				return
			}
			log.Printf("Websocket error: %s", err)
			return
		}
	}
}

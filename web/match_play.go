package web

import (
	"net/http"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"log"
	"github.com/kennhung/ftcScoring/play"
	"fmt"
	"io"
	"github.com/mitchellh/mapstructure"
	"github.com/kennhung/ftcScoring/model"
	"github.com/gorilla/mux"
	"strconv"
)

type MatchTimeMessage struct {
	MatchState   int
	MatchTimeSec int
}

// Global var to hold the current active tournament so that its matches are displayed by default.
var currentMatchType string

func (web *Web) matchPlayHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)

	template.Match_Play("test", buffer)
	w.Write(buffer.Bytes())
}

func (web *Web) matchPlayWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	websocket, err := NewWebsocket(w, r)

	if err != nil {
		handleWebErr(w, err)
		return
	}

	defer websocket.Close()

	matchTimeListener := web.arena.MatchTimeChannel.Listen()
	defer close(matchTimeListener)
	audienceDisplayListener := web.arena.AudienceDisplayChannel.Listen()
	defer close(audienceDisplayListener)
	scoringStatusListener := web.arena.ScoringStatusChannel.Listen()
	defer close(scoringStatusListener)

	var data interface{}

	var MatchStatus = struct {
		MatchState    int
		Teams         map[string]*model.Team
		CanStartMatch bool
	}{web.arena.MatchState, web.arena.Teams, web.arena.CheckCanStartMatch() == nil}
	err = websocket.Write("status", MatchStatus)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}
	err = websocket.Write("matchTiming", play.MatchTiming)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}
	data = MatchTimeMessage{web.arena.MatchState, int(web.arena.MatchTimeSec())}
	err = websocket.Write("matchTime", data)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}

	err = websocket.Write("setAudienceDisplay", web.arena.AudienceDisplayScreen)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}

	go func() {
		for {
			var messageType string
			var message interface{}
			select {
			case matchTimeSec, ok := <-matchTimeListener:
				if !ok {
					return
				}
				messageType = "matchTime"
				message = MatchTimeMessage{web.arena.MatchState, matchTimeSec.(int)}
			}

			err = websocket.Write(messageType, message)
			if err != nil {
				// The client has probably closed the connection; nothing to do here.
				return
			}
		}
	}()

	for {
		messageType, data, err := websocket.Read()
		if err != nil {
			if err == io.EOF {
				// Client has closed the connection; nothing to do here.
				return
			}
			log.Printf("Websocket error: %s", err)
			return
		}

		switch messageType {
		case "substituteTeam":
			args := struct {
				Team     int
				Position string
			}{}
			err = mapstructure.Decode(data, &args)
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			err = web.arena.SubstituteTeam(args.Team, args.Position)
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
		case "toggleBypass":

		case "startMatch":
			args := struct {
				MuteMatchSounds bool
			}{}
			err = mapstructure.Decode(data, &args)
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			web.arena.MuteMatchSounds = args.MuteMatchSounds
			err = web.arena.StartMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
		case "abortMatch":
			err = web.arena.AbortMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
		case "pauseMatch":
			err = web.arena.PauseMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
		case "resumeMatch":
			err = web.arena.ResumeMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
		case "commitResults":
			//err = web.commitCurrentMatchScore()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			err = web.arena.ResetMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			err = web.arena.LoadNextMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			err = websocket.Write("reload", nil)
			if err != nil {
				log.Printf("Websocket error: %s", err)
				return
			}
			continue // Skip sending the status update, as the client is about to terminate and reload.
		case "discardResults":
			err = web.arena.ResetMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			err = web.arena.LoadNextMatch()
			if err != nil {
				websocket.WriteError(err.Error())
				continue
			}
			err = websocket.Write("reload", nil)
			if err != nil {
				log.Printf("Websocket error: %s", err)
				return
			}
			continue // Skip sending the status update, as the client is about to terminate and reload.
		case "setAudienceDisplay":
			screen, ok := data.(string)
			if !ok {
				websocket.WriteError(fmt.Sprintf("Failed to parse '%s' message.", messageType))
				continue
			}
			web.arena.AudienceDisplayScreen = screen
			web.arena.AudienceDisplayChannel.Notify(nil)
			continue
		default:
			websocket.WriteError(fmt.Sprintf("Invalid message type '%s'.", messageType))
			continue
		}

		// Send out the status again after handling the command, as it most likely changed as a result.
		var arenaStatus = struct {
			Teams         map[string]*model.Team
			MatchState    int
			CanStartMatch bool
			IsPaused      bool
		}{web.arena.Teams, web.arena.MatchState, web.arena.CheckCanStartMatch() == nil,web.arena.MatchPaused}

		err = websocket.Write("status", arenaStatus)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			return
		}
	}
}

// Loads the given match onto the arena in preparation for playing it.
func (web *Web) matchPlayLoadHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	matchId, _ := strconv.Atoi(vars["matchId"])
	var match *model.Match
	var err error
	if matchId == 0 {
		err = web.arena.LoadEmptyMatch()
	} else {
		match, err = web.arena.Database.GetMatchById(matchId)
		if err != nil {
			handleWebErr(w, err)
			return
		}
		if match == nil {
			handleWebErr(w, fmt.Errorf("Invalid match ID %d.", matchId))
			return
		}
		err = web.arena.LoadMatch(match)
	}
	if err != nil {
		handleWebErr(w, err)
		return
	}
	currentMatchType = web.arena.CurrentMatch.Type

	http.Redirect(w, r, "/match/play", 303)
}

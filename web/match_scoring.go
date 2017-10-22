package web

import (
	"bytes"
	"net/http"
	"github.com/kennhung/ftcScoring/webTemplate"
	"github.com/kennhung/ftcScoring/model"
	"log"
	"github.com/kennhung/ftcScoring/play"
	"fmt"
	"io"
	"github.com/kennhung/ftcScoring/arena"
)

func (web *Web) matchScoringHandler(w http.ResponseWriter, r *http.Request) {

	var data [3][]model.Match

	var err error

	data[0], err = web.arena.Database.GetMatchesByType("practice")
	if err != nil{
		handleWebErr(w, err)
		return
	}
	/*
		data[1], err = web.arena.Database.GetMatchesByType("qualification")
		if err != nil{
			handleWebErr(w, err)
			return
		}

		data[2], err = web.arena.Database.GetMatchesByType("elimination")
		if err != nil{
			handleWebErr(w, err)
			return
		}
	*/
	buffer := new(bytes.Buffer)
	template.Match_Scoring(data, buffer)

	w.Write(buffer.Bytes())
}

func (web *Web) matchScoringWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	websocket, err := NewWebsocket(w, r)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}

	defer websocket.Close()
	matchLoadTeamsListener := web.arena.MatchLoadTeamsChannel.Listen()
	defer close(matchLoadTeamsListener)
	matchTimeListener := web.arena.MatchTimeChannel.Listen()
	defer close(matchTimeListener)
	reloadDisplaysListener := web.arena.ReloadDisplaysChannel.Listen()
	defer close(reloadDisplaysListener)

	// Send the various notifications immediately upon connection.
	data := struct {
		RedScore  *play.Score
		BlueScore *play.Score
	}{web.arena.RedScore,web.arena.BlueScore}
	err = websocket.Write("score", data)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	}

	err = websocket.Write("matchTime", MatchTimeMessage{web.arena.MatchState, int(web.arena.MatchTimeSec())})
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
			case _, ok := <-matchLoadTeamsListener:
				if !ok {
					return
				}
				messageType = "reload"
				message = nil
			case matchTimeSec, ok := <-matchTimeListener:
				if !ok {
					return
				}
				messageType = "matchTime"
				message = MatchTimeMessage{web.arena.MatchState, matchTimeSec.(int)}
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
		case "commitMatch":
			if web.arena.MatchState != arena.PostMatch {
				// Don't allow committing the score until the match is over.
				websocket.WriteError("Cannot commit score: Match is not over.")
				continue
			}

			web.arena.ScoringStatusChannel.Notify(nil)
		default:
			websocket.WriteError(fmt.Sprintf("Invalid message type '%s'.", messageType))
			continue
		}


		// Send out the score again after handling the command, as it most likely changed as a result.
		data = struct {
			RedScore  *play.Score
			BlueScore *play.Score
		}{}
		err = websocket.Write("score", data)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			return
		}
	}

}

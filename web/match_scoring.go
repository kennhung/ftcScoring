package web

import (
	"bytes"
	"net/http"
	"github.com/kennhung/ftcScoring/webTemplate"
	"github.com/kennhung/ftcScoring/model"
	"log"
	"github.com/kennhung/ftcScoring/play"
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

	currentid := web.arena.CurrentMatch.Id

	matchResult, err := web.arena.Database.GetMatchResultForMatch(currentid)
	matchResult.RedScore.RelicsZ3 = 0;
	if err != nil {
		handleWebErr(w, err);
		return
	}

	websocket, err := NewWebsocket(w, r)
	if err != nil {
		handleWebErr(w, err)
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
	}{}
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

}

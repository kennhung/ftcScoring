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
	"strings"
	"strconv"
)

func (web *Web) matchScoringHandler(w http.ResponseWriter, r *http.Request) {

	var data [3][]model.Match
	var err error

	currentMatch := web.arena.CurrentMatch
	data[0], err = web.arena.Database.GetMatchesByType("practice")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	data[1], err = web.arena.Database.GetMatchesByType("qualification")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	data[2], err = web.arena.Database.GetMatchesByType("elimination")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	buffer := new(bytes.Buffer)
	template.Match_Scoring(data, currentMatch, buffer)

	w.Write(buffer.Bytes())
}

func (web *Web) matchScoringWebsocketHandler(w http.ResponseWriter, r *http.Request) {

	currentMatch := web.arena.CurrentMatch
	currentMatchResult, err := web.arena.Database.GetMatchResultForMatch(currentMatch.Id)
	if err != nil {
		log.Printf("Websocket error: %s", err)
		return
	} else if currentMatchResult == nil {
		//Not Saved yet
		currentMatchResult = model.NewMatchResult()
	}

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
	}{web.arena.RedScore, web.arena.BlueScore}

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
		case "updateRedScore":
			strs := strings.Split(fmt.Sprint(data),":")
			web.updateScore("Red",strs[0],strs[1])

		case "updateBlueScore":
			strs := strings.Split(fmt.Sprint(data),":")
			web.updateScore("Blue",strs[0],strs[1])

		default:
			websocket.WriteError(fmt.Sprintf("Invalid message type '%s'.", messageType))
			continue
		}

		// Send out the score again after handling the command, as it most likely changed as a result.
		data = struct {
			RedScore  *play.Score
			BlueScore *play.Score
		}{web.arena.RedScore,web.arena.BlueScore}
		err = websocket.Write("score", data)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			return
		}
	}

}

func (web *Web) getCurrentMatchResult() *model.MatchResult {
	var RedCards map[string]string
	var BlueCards map[string]string

	if web.arena.Teams["R1"].YellowCard && web.arena.RedScore.Cards["R1"] == "Yellow" {
		RedCards["R1"] = "Red"
	} else if !web.arena.Teams["R1"].YellowCard && web.arena.RedScore.Cards["R1"] == "Yellow" {
		RedCards["R1"] = "Yellow"
	} else if web.arena.RedScore.Cards["R1"] == "Red" {
		RedCards["R1"] = "Red"
	}

	if web.arena.Teams["R2"].YellowCard && web.arena.RedScore.Cards["R2"] == "Yellow" {
		RedCards["R2"] = "Red"
	} else if !web.arena.Teams["R2"].YellowCard && web.arena.RedScore.Cards["R2"] == "Yellow" {
		RedCards["R2"] = "Yellow"
	} else if web.arena.RedScore.Cards["R2"] == "Red" {
		RedCards["R2"] = "Red"
	}
	if web.arena.Teams["B1"].YellowCard && web.arena.BlueScore.Cards["B1"] == "Yellow" {
		BlueCards["B1"] = "Red"
	} else if !web.arena.Teams["B1"].YellowCard && web.arena.BlueScore.Cards["B1"] == "Yellow" {
		BlueCards["B1"] = "Yellow"
	} else if web.arena.BlueScore.Cards["B1"] == "Red" {
		BlueCards["B1"] = "Red"
	}
	if web.arena.Teams["B2"].YellowCard && web.arena.BlueScore.Cards["B2"] == "Yellow" {
		BlueCards["B2"] = "Red"
	} else if !web.arena.Teams["B2"].YellowCard && web.arena.BlueScore.Cards["B2"] == "Yellow" {
		BlueCards["B2"] = "Yellow"
	} else if web.arena.BlueScore.Cards["B2"] == "Red" {
		BlueCards["B2"] = "Red"
	}

	return &model.MatchResult{MatchId: web.arena.CurrentMatch.Id, MatchType: web.arena.CurrentMatch.Type,
		RedScore: web.arena.RedScore, BlueScore: web.arena.BlueScore,
		RedCards: RedCards, BlueCards: BlueCards}
}

func (web *Web) updateScore(color string, field string, data string){
	score := new(play.Score)
	if color == "Red"{//Red
		score = web.arena.RedScore
	} else if color == "Blue"{//Blue
		score = web.arena.BlueScore
	}

	switch field {

	//Autonomous Period
	case "AutoJewels":
		score.AutoJewels,_ =  strconv.Atoi(data)
	case "AutoCryptobox":
		score.AutoCryptobox,_ = strconv.Atoi(data)
	case "CryptoboxKeys":
		score.CryptoboxKeys,_ = strconv.Atoi(data)
	case "RobotInSafeZone":
		score.RobotInSafeZone,_ = strconv.Atoi(data)

	//Driver-Controlled Period
	case "Glyphs":
		score.Glyphs,_ = strconv.Atoi(data)
	case "ComRows":
		score.ComRows,_ = strconv.Atoi(data)
	case "ComColumns":
		score.ComColumns,_ = strconv.Atoi(data)
	case "ComCiphers":
		score.ComCiphers,_ = strconv.Atoi(data)

	//End Game Period
	case "RelicsZ1":
		score.RelicsZ1,_ = strconv.Atoi(data)
	case "RelicsZ2":
		score.RelicsZ2,_ = strconv.Atoi(data)
	case "RelicsZ3":
		score.RelicsZ3,_ = strconv.Atoi(data)
	case "RelicsUpright":
		score.RelicsUpright,_ = strconv.Atoi(data)
	case "RobotBalanced":
		score.RobotBalanced,_ = strconv.Atoi(data)

	//Penalties
	case "MinorPena":
		score.Penalties[1],_ = strconv.Atoi(data)
	case "MajorPena":
		score.Penalties[0],_ = strconv.Atoi(data)
	}
}
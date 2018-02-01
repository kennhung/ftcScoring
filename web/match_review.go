// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Web routes for editing match results.

package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/kennhung/ftcScoring/model"
	"github.com/kennhung/ftcScoring/webTemplate"
	"bytes"
)

type MatchReviewListItem struct {
	Id          int
	DisplayName string
	Time        string
	RedTeams    []int
	BlueTeams   []int
	RedScore    int
	BlueScore   int
	ColorClass  string
}

// Shows the match review interface.
func (web *Web) matchReviewHandler(w http.ResponseWriter, r *http.Request) {

	practiceMatches, err := web.arena.Database.GetMatchesByType("practice")
	if err != nil {
		handleWebErr(w, err)
		return
	}
	qualificationMatches, err := web.arena.Database.GetMatchesByType("qualification")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	eliminationMatches, err := web.arena.Database.GetMatchesByType("elimination")
	if err != nil {
		handleWebErr(w, err)
		return
	}
	matchesByType := [3][]model.Match{practiceMatches,qualificationMatches,eliminationMatches}

	buffer := new(bytes.Buffer)

	template.Match_Review(web.arena.EventSettings, matchesByType, web.arena.CurrentMatch, buffer)
	w.Write(buffer.Bytes())
}

// Shows the page to edit the results for a match.
func (web *Web) matchReviewEditGetHandler(w http.ResponseWriter, r *http.Request) { /*
	match, matchResult, _, err := web.getMatchResultFromRequest(r)
	if err != nil {
		handleWebErr(w, err)
		return
	}

	matchResultJson, err := matchResult.Serialize()
	if err != nil {
		handleWebErr(w, err)
		return
	}
	data := struct {
		*model.EventSettings
		Match           *model.Match
		MatchResultJson *model.MatchResultDb
	}{web.arena.EventSettings, match, matchResultJson}

	if err != nil {
		handleWebErr(w, err)
		return
	}*/
}

// Updates the results for a match.
func (web *Web) matchReviewEditPostHandler(w http.ResponseWriter, r *http.Request) {
	match, matchResult, isCurrent, err := web.getMatchResultFromRequest(r)
	if err != nil {
		handleWebErr(w, err)
		return
	}

	r.ParseForm()
	matchResultJson := model.MatchResultDb{Id: matchResult.Id, MatchId: match.Id, PlayNumber: matchResult.PlayNumber,
		MatchType: matchResult.MatchType, RedScoreJson: r.PostFormValue("redScoreJson"),
		BlueScoreJson: r.PostFormValue("blueScoreJson"), RedCardsJson: r.PostFormValue("redCardsJson"),
		BlueCardsJson: r.PostFormValue("blueCardsJson")}

	// Deserialize the JSON using the same mechanism as to store scoring information in the database.
	matchResult, err = matchResultJson.Deserialize()
	if err != nil {
		handleWebErr(w, err)
		return
	}

	if isCurrent {
		// If editing the current match, just save it back to memory.
		/*
		web.arena.RedRealtimeScore.CurrentScore = *matchResult.RedScore
		web.arena.BlueRealtimeScore.CurrentScore = *matchResult.BlueScore
		web.arena.RedRealtimeScore.Cards = matchResult.RedCards
		web.arena.BlueRealtimeScore.Cards = matchResult.BlueCards
*/
		http.Redirect(w, r, "/match_play", 303)
	} else {
		err = web.commitMatchScore(match, matchResult, false)
		if err != nil {
			handleWebErr(w, err)
			return
		}

		http.Redirect(w, r, "/match_review", 303)
	}
}

// Load the match result for the match referenced in the HTTP query string.
func (web *Web) getMatchResultFromRequest(r *http.Request) (*model.Match, *model.MatchResult, bool, error) {
	vars := mux.Vars(r)

	// If editing the current match, get it from memory instead of the DB.
	if vars["matchId"] == "current" {
		return web.arena.CurrentMatch, web.getCurrentMatchResult(), true, nil
	}

	matchId, _ := strconv.Atoi(vars["matchId"])
	match, err := web.arena.Database.GetMatchById(matchId)
	if err != nil {
		return nil, nil, false, err
	}
	if match == nil {
		return nil, nil, false, fmt.Errorf("Error: No such match: %d", matchId)
	}
	matchResult, err := web.arena.Database.GetMatchResultForMatch(matchId)
	if err != nil {
		return nil, nil, false, err
	}
	if matchResult == nil {
		// We're scoring a match that hasn't been played yet, but that's okay.
		matchResult = model.NewMatchResult()
		matchResult.MatchType = match.Type
	}

	return match, matchResult, false, nil
}

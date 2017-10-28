package arena

import (
	"time"
	"github.com/kennhung/ftcScoring/model"
	"github.com/kennhung/ftcScoring/play"
	"fmt"
)

const (
	LoopPeriodMs          = 10
	matchEndScoreDelaySec = 3
	pickupDelay           = 8
)

type Arena struct {
	Database                 *model.Database
	EventSettings            *model.EventSettings
	CurrentMatch             *model.Match
	MatchState               int
	PrevMatchState           int
	LastMatchUpdateTime      time.Time
	MatchRemainTime          *play.MatchTime
	Teams                    map[string]*model.Team
	RedScore                 *play.Score
	BlueScore                *play.Score
	PrevMatchTimeSec         float64
	AudienceDisplayScreen    string
	SavedMatch               *model.Match
	SavedMatchResult         *model.MatchResult
	FieldTestMode            string
	matchAborted             bool
	MatchPaused              bool
	MuteMatchSounds          bool
	MatchStateChannel        *Channel
	MatchTimeChannel         *Channel
	MatchLoadTeamsChannel    *Channel
	ScoringStatusChannel     *Channel
	ScorePostedChannel       *Channel
	AudienceDisplayChannel   *Channel
	PlaySoundChannel         *Channel
	ReloadDisplaysChannel    *Channel
	PitDisplaysChannel       *Channel
	AllianceSelectionChannel *Channel
}

type ArenaStatus struct {
	Teams         map[string]*model.Team
	MatchState    int
	CanStartMatch bool
}

// Progression of match states.
const (
	PreMatch      = 0
	StartMatch    = 1
	AutoPeriod    = 2
	PickupPeriod  = 3
	TeleopPeriod  = 4
	EndgamePeriod = 5
	PostMatch     = 6
)

// Setup a new arena.
func NewArena(dbPath string) (*Arena, error) {
	arena := new(Arena)
	var err error
	arena.Database, err = model.OpenDatabase(dbPath)
	if err != nil {
		return nil, err
	}

	settings, err := arena.Database.GetEventSettings()
	if err != nil {
		return nil, err
	}
	arena.EventSettings = settings

	//Generate teams
	arena.Teams = make(map[string]*model.Team)
	arena.Teams["R1"] = new(model.Team)
	arena.Teams["R2"] = new(model.Team)
	arena.Teams["B1"] = new(model.Team)
	arena.Teams["B2"] = new(model.Team)

	//Generate channels
	arena.AudienceDisplayChannel = NewChannel()
	arena.MatchLoadTeamsChannel = NewChannel()
	arena.MatchStateChannel = NewChannel()
	arena.MatchTimeChannel = NewChannel()
	arena.ScorePostedChannel = NewChannel()
	arena.ScoringStatusChannel = NewChannel()
	arena.ReloadDisplaysChannel = NewChannel()
	arena.PitDisplaysChannel = NewChannel()
	arena.PlaySoundChannel = NewChannel()
	arena.AllianceSelectionChannel = NewChannel()

	// Load empty match first.
	arena.MatchState = PreMatch
	arena.MatchPaused = false
	arena.PrevMatchState = -1
	arena.PrevMatchTimeSec = 0
	arena.MatchRemainTime = new(play.MatchTime)
	arena.MatchResetTimer()
	arena.LoadEmptyMatch()

	// Initialize display parameters.
	arena.AudienceDisplayScreen = "blank"
	arena.SavedMatch = &model.Match{}
	arena.SavedMatchResult = model.NewMatchResult()

	return arena, nil
}

// Loads the event settings when startup or setting change.
func (arena *Arena) LoadSettings() error {
	settings, err := arena.Database.GetEventSettings()
	if err != nil {
		return err
	}
	arena.EventSettings = settings
	return nil
}

// Sets up the arena for the given match.
func (arena *Arena) LoadMatch(match *model.Match) error {
	if arena.MatchState != PreMatch {
		return fmt.Errorf("Cannot load match while there is a match still in progress or with results pending.")
	}

	arena.CurrentMatch = match
	arena.MatchPaused = false
	err := arena.assignTeam(match.Red1, "R1")
	if err != nil {
		return err
	}
	err = arena.assignTeam(match.Red2, "R2")
	if err != nil {
		return err
	}

	err = arena.assignTeam(match.Blue1, "B1")
	if err != nil {
		return err
	}
	err = arena.assignTeam(match.Blue2, "B2")
	if err != nil {
		return err
	}

	arena.MatchResetTimer()

	arena.RedScore = play.NewScore()
	arena.BlueScore = play.NewScore()

	// Notify any listeners about the new match.
	arena.MatchLoadTeamsChannel.Notify(nil)

	return nil
}

func (arena *Arena) LoadEmptyMatch() error {
	return arena.LoadMatch(&model.Match{Type: "empty"})
}

// Loads the first unplayed match of the current match type.
func (arena *Arena) LoadNextMatch() error {
	if arena.CurrentMatch.Type == "empty" {
		return arena.LoadEmptyMatch()
	}

	matches, err := arena.Database.GetMatchesByType(arena.CurrentMatch.Type)
	if err != nil {
		return err
	}
	for _, match := range matches {
		if match.Status != "complete" {
			err = arena.LoadMatch(&match)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

// Assigns the given team to the given station, also substituting it into the match record.
// Usually use when web change the team number.
func (arena *Arena) SubstituteTeam(teamId int, station string) error {
	if arena.CurrentMatch.Type == "qualification" {
		return fmt.Errorf("Can't substitute teams for qualification matches.")
	}
	err := arena.assignTeam(teamId, station)
	if err != nil {
		return err
	}
	switch station {
	case "R1":
		arena.CurrentMatch.Red1 = teamId
	case "R2":
		arena.CurrentMatch.Red2 = teamId
	case "B1":
		arena.CurrentMatch.Blue1 = teamId
	case "B2":
		arena.CurrentMatch.Blue2 = teamId
	}
	arena.MatchLoadTeamsChannel.Notify(nil)
	return nil
}

func (arena *Arena) assignTeam(teamId int, station string) error {
	// Reject invalid station values.
	if _, ok := arena.Teams[station]; !ok {
		return fmt.Errorf("Invalid alliance station '%s'.", station)
	}
	// Leave the station empty if the team number is zero.
	if teamId == 0 {
		arena.Teams[station] = nil
		return nil
	}

	// Load the team model. If it doesn't exist, enable anonymous operation.
	team, err := arena.Database.GetTeamById(teamId)
	if err != nil {
		return err
	}
	if team == nil {
		team = &model.Team{Id: teamId}
	}

	arena.Teams[station] = team
	return nil
}

// Loops indefinitely to track and update the arena components.
func (arena *Arena) Run() {
	// Start other loops in goroutines.

	for {
		arena.Update()
		time.Sleep(time.Millisecond * LoopPeriodMs)
	}
}

// Starts the match if all conditions are met.
func (arena *Arena) StartMatch() error {
	err := arena.CheckCanStartMatch()
	if err == nil {
		// Save the match start time to the database for posterity.
		arena.CurrentMatch.StartedAt = time.Now()
		if arena.CurrentMatch.Type != "empty" {
			arena.Database.SaveMatch(arena.CurrentMatch)
		}
		arena.MatchState = StartMatch
	}

	if !arena.MuteMatchSounds {
		arena.PlaySoundChannel.Notify("match-start")
	}
	return err
}

// Returns nil if the match can be started, and an error otherwise.
func (arena *Arena) CheckCanStartMatch() error {
	if arena.MatchState != PreMatch {
		return fmt.Errorf("Cannot start match while there is a match still in progress or with results pending.")
	}
	return nil
}

// Kills the current match if it is underway.
func (arena *Arena) AbortMatch() error {
	if arena.MatchState == PreMatch || arena.MatchState == PostMatch {
		return fmt.Errorf("Cannot abort match when it is not in progress.")
	}
	arena.MatchState = PostMatch
	arena.matchAborted = true
	arena.AudienceDisplayScreen = "blank"
	arena.AudienceDisplayChannel.Notify(nil)
	if !arena.MuteMatchSounds {
		arena.PlaySoundChannel.Notify("match-abort")
	}
	return nil
}

// Pause the current Match.
func (arena *Arena) PauseMatch() error {
	if arena.MatchState == PreMatch || arena.MatchState == PostMatch {
		return fmt.Errorf("Cannot pause match when it is not in progress.")
	} else if arena.MatchPaused {
		return fmt.Errorf("Match had been paused.")
	}
	arena.MatchPaused = true
	return nil
}

// Resume the current Match.
func (arena *Arena) ResumeMatch() error {
	if arena.MatchState == PreMatch || arena.MatchState == PostMatch {
		return fmt.Errorf("Cannot resume match when it is not in progress.")
	} else if !arena.MatchPaused {
		return fmt.Errorf("Match is not paused.")
	}

	arena.MatchPaused = false
	if !arena.MuteMatchSounds {
		arena.PlaySoundChannel.Notify("match-resume")
	}
	return nil
}

// Clears out the match and resets the arena state unless there is a match underway.
func (arena *Arena) ResetMatch() error {
	if arena.MatchState != PostMatch && arena.MatchState != PreMatch {
		return fmt.Errorf("Cannot reset match while it is in progress.")
	}
	arena.MatchState = PreMatch
	arena.matchAborted = false
	arena.MuteMatchSounds = false
	arena.MatchResetTimer()
	return nil
}

func (arena *Arena) Update() {
	arena.MatchTimingUpdate()
	matchTimeSec := arena.MatchTimeSec()

	// Decide what state the program should run, depending on where we are in the match.
	switch arena.MatchState {
	case PreMatch:
		matchTimeSec = arena.MatchRemainTime.AutoRemainTime + arena.MatchRemainTime.TeleopRemainTime + arena.MatchRemainTime.EndgameRemainTime
		arena.PrevMatchTimeSec = -1
	case StartMatch:
		arena.MatchState = AutoPeriod
		matchTimeSec = arena.MatchRemainTime.AutoRemainTime + arena.MatchRemainTime.TeleopRemainTime + arena.MatchRemainTime.EndgameRemainTime
		arena.PrevMatchTimeSec = -1
		arena.AudienceDisplayScreen = "matchTimer"
	case AutoPeriod:
		if matchTimeSec <= 0 {
			arena.MatchState = PickupPeriod
			if !arena.MuteMatchSounds {
				arena.PlaySoundChannel.Notify("auto-end")
			}
		}

	case PickupPeriod:
		if matchTimeSec <= 0 {
			arena.MatchState = TeleopPeriod
			if !arena.MuteMatchSounds {
				arena.PlaySoundChannel.Notify("match-resume")
			}
		}
	case TeleopPeriod:
		if matchTimeSec <= 0 {
			arena.MatchState = EndgamePeriod
			if !arena.MuteMatchSounds {
				arena.PlaySoundChannel.Notify("match-endgame")
			}
		}
	case EndgamePeriod:
		if matchTimeSec <= 0 {
			arena.MatchState = PostMatch
			if !arena.MuteMatchSounds {
				arena.PlaySoundChannel.Notify("match-end")
			}
			arena.MatchTimeChannel.Notify(int(arena.MatchRemainTime.AutoRemainTime + arena.MatchRemainTime.TeleopRemainTime + arena.MatchRemainTime.EndgameRemainTime))
			go func() {
				// Leave the scores on the screen briefly at the end of the match.
				time.Sleep(time.Second * matchEndScoreDelaySec)
				arena.AudienceDisplayScreen = "blank"
				arena.AudienceDisplayChannel.Notify(nil)
			}()
		}

	}

	if arena.PrevMatchState != arena.MatchState {
		arena.MatchStateChannel.Notify(nil)
		arena.PrevMatchState = arena.MatchState
	}

	if int(matchTimeSec) != int(arena.PrevMatchTimeSec) {
		if arena.MatchState == PickupPeriod {
			arena.MatchTimeChannel.Notify(int(arena.MatchRemainTime.PickupRemainTime))
		} else {
			arena.MatchTimeChannel.Notify(int(arena.MatchRemainTime.AutoRemainTime + arena.MatchRemainTime.TeleopRemainTime + arena.MatchRemainTime.EndgameRemainTime))
		}
	}
	arena.PrevMatchTimeSec = matchTimeSec
}

// Returns the fractional number of seconds current state remain.
func (arena *Arena) MatchTimeSec() float64 {
	if arena.MatchState == PreMatch || arena.MatchState == StartMatch || arena.MatchState == PostMatch {
		return 0
	} else {
		switch arena.MatchState {
		case AutoPeriod:
			return arena.MatchRemainTime.AutoRemainTime
		case PickupPeriod:
			return arena.MatchRemainTime.PickupRemainTime
		case TeleopPeriod:
			return arena.MatchRemainTime.TeleopRemainTime
		case EndgamePeriod:
			return arena.MatchRemainTime.EndgameRemainTime
		}
	}
	return -1
}

func (arena *Arena) MatchTimingUpdate() {
	if arena.MatchState == PreMatch || arena.MatchState == StartMatch || arena.MatchState == PostMatch || arena.MatchPaused {
		arena.LastMatchUpdateTime = time.Now()
		return
	} else {
		switch arena.MatchState {
		case AutoPeriod:
			arena.MatchRemainTime.AutoRemainTime -= time.Since(arena.LastMatchUpdateTime).Seconds()
		case PickupPeriod:
			arena.MatchRemainTime.PickupRemainTime -= time.Since(arena.LastMatchUpdateTime).Seconds()
		case TeleopPeriod:
			arena.MatchRemainTime.TeleopRemainTime -= time.Since(arena.LastMatchUpdateTime).Seconds()
		case EndgamePeriod:
			arena.MatchRemainTime.EndgameRemainTime -= time.Since(arena.LastMatchUpdateTime).Seconds()
		}
	}
	arena.LastMatchUpdateTime = time.Now()
}

func (arena *Arena) MatchResetTimer() {
	arena.MatchRemainTime.AutoRemainTime = float64(play.MatchTiming.AutoDurationSec)
	arena.MatchRemainTime.PickupRemainTime = float64(play.MatchTiming.PickupDurationSec)
	arena.MatchRemainTime.TeleopRemainTime = float64(play.MatchTiming.TeleopDurationSec)
	arena.MatchRemainTime.EndgameRemainTime = float64(play.MatchTiming.EndgameTimeLeftSec)

}

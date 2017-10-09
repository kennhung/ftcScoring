package arena

import (
	"time"
	"github.com/kennhung/ftcScoring/model"
)

const (
	LoopPeriodMs     = 10
	matchEndScoreDelaySec = 3
)

type Arena struct {
	Database                       *model.Database
	EventSettings                  *model.EventSettings
	CurrentMatch                   *model.Match
	MatchState                     int
	lastMatchState                 int
	MatchStartTime                 time.Time
	LastMatchTimeSec               float64
	FieldReset                     bool
	AudienceDisplayScreen          string
	SavedMatch                     *model.Match
	SavedMatchResult               *model.MatchResult
	FieldTestMode                  string
	matchAborted                   bool
}

func NewArena(dbPath string) (*Arena, error) {
	arena := new(Arena)
	var err error
	arena.Database, err = model.OpenDatabase(dbPath)
	if err != nil {
		return nil, err
	}

	settings, err := arena.Database.GetEventSettings()
	if err != nil {
		return nil,err
	}
	arena.EventSettings = settings

	return arena, nil
}

// Loops indefinitely to track and update the arena components.
func (arena *Arena) Run() {
	// Start other loops in goroutines.

	for {
		arena.Update()
		time.Sleep(time.Millisecond * LoopPeriodMs)
	}
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

func (arena *Arena) Update() {/*
	// Decide what state the robots need to be in, depending on where we are in the match.
	auto := false
	enabled := false
	sendDsPacket := false
	matchTimeSec := arena.MatchTimeSec()
	switch arena.MatchState {
	case PreMatch:
		auto = true
		enabled = false
	case StartMatch:
		arena.MatchState = AutoPeriod
		arena.MatchStartTime = time.Now()
		arena.LastMatchTimeSec = -1
		auto = true
		enabled = true
		sendDsPacket = true
		arena.AudienceDisplayScreen = "match"
		arena.AudienceDisplayNotifier.Notify(nil)
		if !arena.MuteMatchSounds {
			arena.PlaySoundNotifier.Notify("match-start")
		}
		arena.FieldTestMode = ""
		arena.Plc.ResetCounts()
	case AutoPeriod:
		auto = true
		enabled = true
		if matchTimeSec >= float64(game.MatchTiming.AutoDurationSec) {
			arena.MatchState = PausePeriod
			auto = false
			enabled = false
			sendDsPacket = true
			if !arena.MuteMatchSounds {
				arena.PlaySoundNotifier.Notify("match-end")
			}
		}
	case PausePeriod:
		auto = false
		enabled = false
		if matchTimeSec >= float64(game.MatchTiming.AutoDurationSec+game.MatchTiming.PauseDurationSec) {
			arena.MatchState = TeleopPeriod
			auto = false
			enabled = true
			sendDsPacket = true
			if !arena.MuteMatchSounds {
				arena.PlaySoundNotifier.Notify("match-resume")
			}
		}
	case TeleopPeriod:
		auto = false
		enabled = true
		if matchTimeSec >= float64(game.MatchTiming.AutoDurationSec+game.MatchTiming.PauseDurationSec+
			game.MatchTiming.TeleopDurationSec-game.MatchTiming.EndgameTimeLeftSec) {
			arena.MatchState = EndgamePeriod
			sendDsPacket = false
			if !arena.MuteMatchSounds {
				arena.PlaySoundNotifier.Notify("match-endgame")
			}
		}
	case EndgamePeriod:
		auto = false
		enabled = true
		if matchTimeSec >= float64(game.MatchTiming.AutoDurationSec+game.MatchTiming.PauseDurationSec+
			game.MatchTiming.TeleopDurationSec) {
			arena.MatchState = PostMatch
			auto = false
			enabled = false
			sendDsPacket = true
			go func() {
				// Leave the scores on the screen briefly at the end of the match.
				time.Sleep(time.Second * matchEndScoreDwellSec)
				arena.AudienceDisplayScreen = "blank"
				arena.AudienceDisplayNotifier.Notify(nil)
				arena.AllianceStationDisplayScreen = "logo"
				arena.AllianceStationDisplayNotifier.Notify(nil)
			}()
			if !arena.MuteMatchSounds {
				arena.PlaySoundNotifier.Notify("match-end")
			}
		}
	}

	// Send a notification if the match state has changed.
	if arena.MatchState != arena.lastMatchState {
		arena.matchStateNotifier.Notify(arena.MatchState)
	}
	arena.lastMatchState = arena.MatchState

	// Send a match tick notification if passing an integer second threshold.
	if int(matchTimeSec) != int(arena.LastMatchTimeSec) {
		arena.MatchTimeNotifier.Notify(int(matchTimeSec))
	}
	arena.LastMatchTimeSec = matchTimeSec

	// Send a packet if at a period transition point or if it's been long enough since the last one.
	if sendDsPacket || time.Since(arena.lastDsPacketTime).Seconds()*1000 >= dsPacketPeriodMs {
		arena.sendDsPacket(auto, enabled)
		arena.RobotStatusNotifier.Notify(nil)
	}*/
}
package play

import (
	"time"
)

type MatchTime struct {
	AutoRemainTime float64
	PickupRemainTime float64
	TeleopRemainTime float64
	EndgameRemainTime float64
}

var MatchTiming = struct {
	AutoDurationSec    int
	PickupDurationSec   int
	TeleopDurationSec  int
	EndgameTimeLeftSec int
}{30, 8, 90, 30}

func GetMatchEndTime(matchStartTime time.Time) time.Time {
	return matchStartTime.Add(time.Duration(MatchTiming.AutoDurationSec+MatchTiming.PickupDurationSec+
		MatchTiming.TeleopDurationSec) * time.Second)
}
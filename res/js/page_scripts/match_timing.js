var matchStates = {
    0: "PRE_MATCH",
    1: "START_MATCH",
    2: "AUTO_PERIOD",
    3: "PAUSE_PERIOD",
    4: "TELEOP_PERIOD",
    5: "ENDGAME_PERIOD",
    6: "POST_MATCH"
};
var matchTiming;

// Handles a websocket message containing the length of each period in the match.
var handleMatchTiming = function (data) {
    matchTiming = data;
};

// Converts the raw match state and time into a human-readable state and per-period time. Calls the provided
// callback with the result.
var translateMatchTime = function (data, callback) {
    var matchStateText;
    switch (matchStates[data.MatchState]) {
        case "PRE_MATCH":
            matchStateText = "PRE-MATCH";
            break;
        case "START_MATCH":
        case "AUTO_PERIOD":
            matchStateText = "AUTONOMOUS";
            break;
        case "PAUSE_PERIOD":
            matchStateText = "Pickup";
            break;
        case "TELEOP_PERIOD":
            matchStateText = "TELEOPERATED";
            break;
        case "ENDGAME_PERIOD":
            matchStateText = "TELEOPERATED-ENDGAME";
            break;
        case "POST_MATCH":
            matchStateText = "POST-MATCH";
            break;
    }
    var matchTimeText;
    if(data.MatchTimeSec%60<10){
        matchTimeText = Math.floor(data.MatchTimeSec/60) + ":0" + data.MatchTimeSec%60
    }else{
        matchTimeText = Math.floor(data.MatchTimeSec/60) + ":" + data.MatchTimeSec%60
    }
    callback(matchStates[data.MatchState], matchStateText, matchTimeText);
};

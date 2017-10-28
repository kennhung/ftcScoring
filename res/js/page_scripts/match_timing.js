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

var matchLength = 150;

// Handles a websocket message containing the length of each period in the match.
var handleMatchTiming = function (data) {
    matchTiming = data;
};

// Converts the raw match state and time into a human-readable state and per-period time. Calls the provided
// callback with the result.
var translateMatchTime = function (data, callback) {
    var matchStateText;
    var barWidth = "width: ";
    var barClass = "progress-bar "
    switch (matchStates[data.MatchState]) {
        case "PRE_MATCH":
            matchStateText = "PRE-MATCH";
            break;
        case "START_MATCH":
        case "AUTO_PERIOD":
            matchStateText = "AUTONOMOUS";
            barClass += "bg-success";
            break;
        case "PAUSE_PERIOD":
            matchStateText = "Pickup";
            barClass += "bg-danger";
            break;
        case "TELEOP_PERIOD":
            matchStateText = "TELEOPERATED";
            barClass += "bg-success";
            break;
        case "ENDGAME_PERIOD":
            matchStateText = "TELEOPERATED-ENDGAME";
            barClass += "bg-warning";
            barClass
            break;
        case "POST_MATCH":
            matchStateText = "POST-MATCH";
            barClass += "bg-danger";
            break;
    }

    barWidth += data.MatchTimeSec / matchLength * 100;
    barWidth += "%;";
    barWidth += " margin-left:";
    barWidth += 100 - (data.MatchTimeSec / matchLength * 100);
    barWidth += "%; height: 30px;"


    var matchTimeText;
    if (data.MatchTimeSec % 60 < 10) {
        matchTimeText = Math.floor(data.MatchTimeSec / 60) + ":0" + data.MatchTimeSec % 60
    } else {
        matchTimeText = Math.floor(data.MatchTimeSec / 60) + ":" + data.MatchTimeSec % 60
    }
    callback(matchStates[data.MatchState], matchStateText, matchTimeText, barWidth, barClass);
};

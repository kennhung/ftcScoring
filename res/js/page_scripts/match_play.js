var websocket;
var scoreIsReady;
var isPaused;

// Sends a websocket message to load a team into an alliance station.
var substituteTeam = function (team, position) {
    websocket.send("substituteTeam", {team: parseInt(team), position: position})
};

var togglePause = function () {
    if (isPaused) {
        resumeMatch();
    }
    else {
        pauseMatch();
    }
}

// Sends a websocket message to toggle the bypass status for an alliance station.
var toggleBypass = function (station) {
    websocket.send("toggleBypass", station);
};

// Sends a websocket message to start the match.
var startMatch = function () {
    websocket.send("startMatch", {muteMatchSounds: $("#muteMatchSounds").prop("checked")});
};

// Sends a websocket message to abort the match.
var abortMatch = function () {
    websocket.send("abortMatch");
};

// Send a websocket message to pause the match.
var pauseMatch = function () {
    websocket.send("pauseMatch");
}

// Send a websocket message to resume the match.
var resumeMatch = function () {
    websocket.send("resumeMatch");
}


// Sends a websocket message to commit the match score and load the next match.
var commitResults = function () {
    websocket.send("commitResults");
};

// Sends a websocket message to discard the match score and load the next match.
var discardResults = function () {
    websocket.send("discardResults");
};

// Sends a websocket message to change what the audience display is showing.
var setAudienceDisplay = function () {
    websocket.send("setAudienceDisplay", $("input[name=audienceDisplay]:checked").val());
};

// Sends a websocket message to change what the alliance station display is showing.
var setAllianceStationDisplay = function () {
    websocket.send("setAllianceStationDisplay", $("input[name=allianceStationDisplay]:checked").val());
};

var confirmCommit = function (isReplay) {
    if (isReplay || !scoreIsReady) {
        // Show the appropriate message(s) in the confirmation dialog.
        $("#confirmCommitReplay").css("display", isReplay ? "block" : "none");
        $("#confirmCommitNotReady").css("display", scoreIsReady ? "none" : "block");
        $("#confirmCommitResults").modal("show");
    } else {
        commitResults();
    }
};

// Handles a websocket message to update the team connection status.
var handleStatus = function (data) {

    // Update the team status view.
    $.each(data.Teams, function (index, team) {

    });

    isPaused = data.IsPaused;

    if (isPaused) {
        $("#togglePause").text("Resume Match")
    }
    else {
        $("#togglePause").text("Pause Match")
    }

    // Enable/disable the buttons based on the current match state.
    switch (matchStates[data.MatchState]) {
        case "PRE_MATCH":
            $("#startMatch").prop("disabled", !data.CanStartMatch);
            $("#togglePause").prop("disabled", true);
            $("#abortMatch").prop("disabled", true);
            $("#commitResults").prop("disabled", true);
            $("#discardResults").prop("disabled", true);
            $("#editResults").prop("disabled", true);
            break;
        case "START_MATCH":
        case "AUTO_PERIOD":
        case "PAUSE_PERIOD":
        case "TELEOP_PERIOD":
        case "ENDGAME_PERIOD":
            $("#startMatch").prop("disabled", true);
            $("#togglePause").prop("disabled", false);
            $("#abortMatch").prop("disabled", false);
            $("#commitResults").prop("disabled", true);
            $("#discardResults").prop("disabled", true);
            $("#editResults").prop("disabled", true);
            break;
        case "POST_MATCH":
            $("#startMatch").prop("disabled", true);
            $("#togglePause").prop("disabled", true);
            $("#abortMatch").prop("disabled", true);
            $("#commitResults").prop("disabled", false);
            $("#discardResults").prop("disabled", false);
            $("#editResults").prop("disabled", false);
            break;
    }
};

// Handles a websocket message to update the match time countdown.
var handleMatchTime = function (data) {
    translateMatchTime(data, function (matchState, matchStateText, countdownSec,barWidth,barClass) {
        $("#matchState").text(matchStateText);
        $("#matchTime span").text(countdownSec);

        switch(data.MatchState){
            case 0://Pre
                $(".btn-load").removeClass("disabled");
                $("#timerBar").attr("style",barWidth);
                break;
            case 1:
            case 2:
                $(".btn-load").addClass("disabled");
                $("#timerBar").attr("style",barWidth);
                break;
            case 3://Pause
                $(".btn-load").addClass("disabled");
                break;
            case 4:
            case 5:
                $(".btn-load").addClass("disabled");
                $("#timerBar").attr("style",barWidth);
                break;
            case 6://POST
                $("#timerBar").attr("style","width:100%; height: 30px;");
                $(".btn-load").addClass("disabled");
                break;
        }

        $("#timerBar").attr("class",barClass);
    });
};

// Handles a websocket message to update the match score.
var handleRealtimeScore = function (data) {
    $("#redScore").text(data.RedScore);
    $("#blueScore").text(data.BlueScore);
};

// Handles a websocket message to update the audience display screen selector.
var handleSetAudienceDisplay = function (data) {
    $("input[name=audienceDisplay]:checked").prop("checked", false);
    $("input[name=audienceDisplay][value=" + data + "]").prop("checked", true);
};

// Handles a websocket message to signal whether the referee and scorers have committed after the match.
var handleScoringStatus = function (data) {
    scoreIsReady = data.RefereeScoreReady && data.RedScoreReady && data.BlueScoreReady;
    $("#refereeScoreStatus").attr("data-ready", data.RefereeScoreReady);
    $("#redScoreStatus").attr("data-ready", data.RedScoreReady);
    $("#blueScoreStatus").attr("data-ready", data.BlueScoreReady);
};

// Handles a websocket message to update the alliance station display screen selector.
var handleSetAllianceStationDisplay = function (data) {
    $("input[name=allianceStationDisplay]:checked").prop("checked", false);
    $("input[name=allianceStationDisplay][value=" + data + "]").prop("checked", true);
};

$(function () {
    // Activate tooltips above the status headers.
    $("[data-toggle=tooltip]").tooltip({"placement": "top"});

    // Set up the websocket..
    websocket = new ScoringWebsocket("/match/play/socket", {
        status: function (event) {
            handleStatus(event.data);
        },
        matchTiming: function (event) {
            handleMatchTiming(event.data);
        },
        matchTime: function (event) {
            handleMatchTime(event.data);
        },
        realtimeScore: function (event) {
            handleRealtimeScore(event.data);
        },
        setAudienceDisplay: function (event) {
            handleSetAudienceDisplay(event.data);
        },
        scoringStatus: function (event) {
            handleScoringStatus(event.data);
        },
        setAllianceStationDisplay: function (event) {
            handleSetAllianceStationDisplay(event.data);
        }
    });
});

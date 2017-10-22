var websocket;
var scoreCommitted = false;

function InitMaxandMin() {
    $(".numInput").attr("min", 0);
    $(".numInput").attr('value',0);
    $("#AutoJewels").attr("max", 4);
    $("#AutoCryptobox").attr("max", 48);
    $("#CryptoboxKeys").attr("max", 2);
    $("#RobotInSafeZone").attr("max", 2);
}

var handleScore = function(data) {
    // Update autonomous period values.
    var score = data.Score.CurrentScore;
    $("#autoMobility").text(score.AutoMobility);

    // Update component visibility.
    if (!data.AutoCommitted) {
        $("#autoScoring").fadeTo(0, 1);
        $("#teleopScoring").hide();
        $("#waitingMessage").hide();
        scoreCommitted = false;
    } else if (!data.Score.TeleopCommitted) {
        $("#autoScoring").fadeTo(0, 0.25);
        $("#teleopScoring").show();
        $("#waitingMessage").hide();
        scoreCommitted = false;
    } else {
        $("#autoScoring").hide();
        $("#teleopScoring").hide();
        $("#commitMatchScore").hide();
        $("#waitingMessage").show();
        scoreCommitted = true;
    }
};

// Handles a websocket message to update the match status.
var handleMatchTime = function(data) {
    if (matchStates[data.MatchState] == "POST_MATCH" && !scoreCommitted) {
        $("#scoringCard").show();
    } else {
        $("#scoringCard").hide();
    }
};


var commitMatchScore = function() {
    websocket.send("commitMatch");
};

$(function() {
    // Set up the websocket back to the server.
    websocket = new ScoringWebsocket("/displays/scoring/websocket", {
        score: function(event) { handleScore(event.data); },
        matchTime: function(event) { handleMatchTime(event.data); }

    });
});

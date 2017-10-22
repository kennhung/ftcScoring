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

    //Update RedScore
    var RedScore = data.RedScore;
    $("#redScoreForm #AutoJewels").val(RedScore.AutoJewels)
    $("#redScoreForm #AutoCryptobox").val(RedScore.AutoCryptobox)
    $("#redScoreForm #CryptoboxKeys").val(RedScore.CryptoboxKeys)
    $("#redScoreForm #RobotInSafeZone").val(RedScore.RobotInSafeZone)

    scoreCommitted = false;
};

// Handles a websocket message to update the match status.
var handleMatchTime = function(data) {
    if (!scoreCommitted) {
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
    websocket = new ScoringWebsocket("/match/scoring/websocket", {
        score: function(event) { handleScore(event.data); },
        matchTime: function(event) { handleMatchTime(event.data); }

    });
});

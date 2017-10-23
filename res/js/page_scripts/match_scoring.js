var websocket;
var scoreCommitted = false;

function InitMaxandMin() {
    $(".numInput").attr("min", 0);
    $(".numInput").attr('value',0);

    //Autonomous Period
    $("#AutoJewels").attr("max", 4);
    $("#AutoCryptobox").attr("max", 48);
    $("#CryptoboxKeys").attr("max", 2);
    $("#RobotInSafeZone").attr("max", 2);

    //Driver-Controlled Period
    $("#Glyphs").attr("max",48)
    $("#ComRows").attr("max",8)
    $("#ComColumns").attr("max",6)
    $("#ComCiphers").attr("max",2)

    //End Game Period
    $("#RelicsZ1").attr("max",2)
    $("#RelicsZ2").attr("max",2)
    $("#RelicsZ3").attr("max",2)
    $("#RelicsUpright").attr("max",2)
    $("#RobotBalanced").attr("max",2)

}

var handleScore = function(data) {

    //Update RedScore
    var RedScore = data.RedScore;
    //Autonomous Period
    $("#redScoreForm #AutoJewels").val(RedScore.AutoJewels)
    $("#redScoreForm #AutoCryptobox").val(RedScore.AutoCryptobox)
    $("#redScoreForm #CryptoboxKeys").val(RedScore.CryptoboxKeys)
    $("#redScoreForm #RobotInSafeZone").val(RedScore.RobotInSafeZone)
    //Driver-Controlled Period
    $("#redScoreForm #Glyphs").val(RedScore.Glyphs)
    $("#redScoreForm #ComRows").val(RedScore.ComRows)
    $("#redScoreForm #ComColumns").val(RedScore.ComColumns)
    $("#redScoreForm #ComCiphers").val(RedScore.ComCiphers)
    //End Game Period
    $("#redScoreForm #RelicsZ1").val(RedScore.RelicsZ1)
    $("#redScoreForm #RelicsZ2").val(RedScore.RelicsZ2)
    $("#redScoreForm #RelicsZ3").val(RedScore.RelicsZ3)
    $("#redScoreForm #RelicsUpright").val(RedScore.RelicsUpright)
    $("#redScoreForm #RobotBalanced").val(RedScore.RobotBalanced)
    //Penalties
    $("#redScoreForm #MinorPena").val(RedScore.Penalties[1])
    $("#redScoreForm #MajorPena").val(RedScore.Penalties[0])
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

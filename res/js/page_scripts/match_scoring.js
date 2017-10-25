var websocket;
var scoreCommitted = false;

function InitMaxandMin() {
    $(".numInput").attr("min", 0);
    $(".numInput").attr('value',0);

    //Red
    //Autonomous Period
    $("#redScoreForm #AutoJewels").attr("max", 4);
    $("#redScoreForm #AutoCryptobox").attr("max", 48);
    $("#redScoreForm #CryptoboxKeys").attr("max", 2);
    $("#redScoreForm #RobotInSafeZone").attr("max", 2);

    //Driver-Controlled Period
    $("#redScoreForm #Glyphs").attr("max",48)
    $("#redScoreForm #ComRows").attr("max",8)
    $("#redScoreForm #ComColumns").attr("max",6)
    $("#redScoreForm #ComCiphers").attr("max",2)

    //End Game Period
    $("#redScoreForm #RelicsZ1").attr("max",2)
    $("#redScoreForm #RelicsZ2").attr("max",2)
    $("#redScoreForm #RelicsZ3").attr("max",2)
    $("#redScoreForm #RelicsUpright").attr("max",2)
    $("#redScoreForm #RobotBalanced").attr("max",2)

    //Blue
    //Autonomous Period
    $("#blueScoreForm #AutoJewels").attr("max", 4);
    $("#blueScoreForm #AutoCryptobox").attr("max", 48);
    $("#blueScoreForm #CryptoboxKeys").attr("max", 2);
    $("#blueScoreForm #RobotInSafeZone").attr("max", 2);

    //Driver-Controlled Period
    $("#blueScoreForm #Glyphs").attr("max",48)
    $("#blueScoreForm #ComRows").attr("max",8)
    $("#blueScoreForm #ComColumns").attr("max",6)
    $("#blueScoreForm #ComCiphers").attr("max",2)

    //End Game Period
    $("#blueScoreForm #RelicsZ1").attr("max",2)
    $("#blueScoreForm #RelicsZ2").attr("max",2)
    $("#blueScoreForm #RelicsZ3").attr("max",2)
    $("#blueScoreForm #RelicsUpright").attr("max",2)
    $("#blueScoreForm #RobotBalanced").attr("max",2)

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

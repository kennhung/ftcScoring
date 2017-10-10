var websocket;

var handletime = function(data){
    $("#evdate").text(data.Date)
    $("#evname").text(data.Name)
    $("#evreg").text(data.Region)
    $("#evtype").text(data.Type)
}


$(function() {
    websocket = new ScoringWebsocket("/match/play/socket",{
        test: function (event) {handletime(event.data);}
    });
});
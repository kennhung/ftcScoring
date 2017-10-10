var ScoringWebsocket = function(path, events) {
    var me = this;

    //Formatting url
    var protocol = "ws://";
    if (window.location.protocol == "https:") {
        protocol = "wss://";
    }
    var url = protocol + window.location.hostname;
    if (window.location.port != "") {
        url += ":" + window.location.port;
    }
    url += path;

    // Show error msg
    if (!events.hasOwnProperty("error")) {
        events.error = function(event) {
            // Data is just an error string.
            console.log(event.data);
        }
    }

    // Show alert
    events.dialog = function(event) {
        alert(event.data);
    }

    // Force reload
    events.reload = function(event) {
        location.reload();
    };

    this.connect = function() {
        this.websocket = $jQuery_1_11_0.websocket(url, {
            open: function() {
                console.log("Websocket connected to the server at " + url + ".")
            },
            close: function() {
                console.log("Websocket lost connection to the server. Reconnecting in 3 seconds...");
                setTimeout(me.connect, 3000);
            },
            events: events
        });
    };

    this.send = function(type, data) {
        this.websocket.send(type, data);
    };

    // connect to socket server
    this.connect();
};

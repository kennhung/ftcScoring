// Code generated by hero.
// source: C:\work\src\github.com\kennhung\ftcScoring\webTemplate\match_play.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"fmt"

	"github.com/kennhung/ftcScoring/model"
	"github.com/shiyanhui/hero"
)

func Match_Play(allMatchs [3][]model.Match, currentMatch *model.Match, paused bool, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <!-- Bootstrap CSS -->
    <link href="/res/css/bootstrap.min.css" rel="stylesheet">
    <link href="/res/open-iconic/font/css/open-iconic-bootstrap.min.css" rel="stylesheet">
    <link href="/res/css/ftcScoring.css" rel="stylesheet">
    <!-- JS Scripts -->

    <!-- load jQuery 1.11.0 -->
    <script src="/res/js/jquery_v1.11.0.min.js"></script>
    <script src="/res/js/jquery.websocket-0.0.1.js"></script>
    <script src="/res/js/jquery.json-2.4.min.js"></script>
    <script src="/res/js/page_scripts/ScoringWebsocket.js"></script>
    <script type="text/javascript">
        $jQuery_1_11_0 = $.noConflict(true);
    </script>

    <script src="/res/js/jquery-3.2.1.slim.min.js"></script>
    <script src="/res/js/popper.min.js"></script>
    <script src="/res/js/bootstrap.min.js"></script>
    <script src="/res/js/bootstrap-number-input.js"></script>
    <!-- Scoring System Script -->
    <script src="/res/js/ftcScoring.js"></script>

    `)
	buffer.WriteString(`
<script src="/res/js/page_scripts/match_play.js"></script>
<script src="/res/js/page_scripts/match_timing.js"></script>
`)

	buffer.WriteString(`

    <title>`)
	buffer.WriteString(`
Match Play
`)

	buffer.WriteString(`- FTC Scoring</title>
</head>
<body style="padding-top: 70px;">
<div style="margin-bottom: 20px;">
    <nav class="navbar fixed-top  navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <a class="navbar-brand" href="/">FTC Scoring</a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
                    aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav mr-auto" id="navbar-menu">
                    <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownSetup" data-toggle="dropdown"
                           aria-haspopup="true" aria-expanded="false">
                            Setup
                        </a>
                        <div class="dropdown-menu">
                            <a class="dropdown-item" href="/setup/settings">Event Settings</a>
                            <a class="dropdown-item" href="/setup/teams">Teams</a>
                            <a class="dropdown-item" href="/setup/schedule">Generate Match</a>
                        </div>
                    </li>
                    <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMatch" data-toggle="dropdown"
                           aria-haspopup="true" aria-expanded="false">
                            Match
                        </a>
                        <div class="dropdown-menu">
                            <a class="dropdown-item" href="/match/play">Play</a>
                            <a class="dropdown-item" href="/match/scoring">Scoring</a>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</div>

<div class="container-fluid">
    `)

	var matchstr = ""
	switch currentMatch.Type {
	case "practice":
		matchstr += fmt.Sprint("P-")
	case "qualification":
		matchstr += fmt.Sprint("Q-")
	case "elimination":
		matchstr += fmt.Sprint("E-")
	case "empty":
		matchstr += fmt.Sprint("Test Match")
	}

	matchstr += fmt.Sprint(currentMatch.DisplayName)

	buffer.WriteString(`

<div class="row">
    <div class="col-lg-1">

    </div>
    <div class="col-lg-3">
        <div class="card bg-light mb-3 border-primary">
            <div class="card-header">
                <ul class="nav nav-tabs card-header-tabs" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link active" role="tab" data-toggle="tab" href="#practice">Practice</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" role="tab" data-toggle="tab" href="#qual">Qualification</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" role="tab" data-toggle="tab" href="#playoff">Playoff</a>
                    </li>
                </ul>
            </div>
            <div class="card-body">
                <div class="tab-content">
                    <div id="practice" class="tab-pane fade show active" role="tabpanel">
                        <h3>Practice</h3>
                        <p>Some content.</p>
                    </div>
                    <div id="qual" class="tab-pane fade" role="tabpanel">
                        <h3>Qualification</h3>
                        <p>Some content in menu 1.</p>
                    </div>
                    <div id="playoff" class="tab-pane fade" role="tabpanel">
                        <h3>Playoff</h3>
                        <p>Some content in menu 2.</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="col-lg">
        <div class="card-deck">
            <div class="card bg-light mb-3 text-center">
                <div class="card-header">
                    Match Status - `)
	hero.EscapeHTML(matchstr, buffer)
	buffer.WriteString(`
                </div>
                <div class="card-body">
                    <div class="progress" style="height: 30px;">
                        <span id="matchTime"
                              style="position:absolute; right:0; left:0; color: #fff; margin: 0; top: 50%; left: 50%;transform: translate(-50%, -50%);"></span>
                        <div class="progress-bar" id="timerBar" role="progressbar" role="progressbar"></div>
                    </div>
                    matchState: <span id="matchState"></span>
                </div>
            </div>
            <div class="card bg-light mb-3 text-center">
                <div class="card-header">
                    Match Control
                </div>
                <div class="card-body">
                    <button class="btn btn-outline-primary mb-1" id="startMatch" onclick="startMatch();">Start Match
                    </button>
                    <button class="btn btn-outline-warning mb-1" id="togglePause" onclick="togglePause();">Pause Match
                    </button>
                    <button class="btn btn-outline-danger mb-1" id="abortMatch" onclick="abortMatch();">AbortMatch
                    </button>
                    <button class="btn btn-outline-warning mb-1" id="discardResults" onclick="discardResults();">Discard
                        Results
                    </button>
                    <button class="btn btn-outline-info mb-1" id="commitResults" onclick="commitResults();">Commit
                        Results
                    </button>
                </div>
            </div>
        </div>
    </div>
    <div class="col-lg-1">

    </div>
</div>

<script>

</script>

`)

	buffer.WriteString(`
</div>
</body>
</html>`)

}

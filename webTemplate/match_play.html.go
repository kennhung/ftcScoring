// Code generated by hero.
// source: C:\work\src\github.com\kennhung\ftcScoring\webTemplate\match_play.html
// DO NOT EDIT!
package template

import "bytes"

func Match_Play(tempstr string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <!-- Bootstrap CSS -->
    <link href="/res/css/bootstrap.min.css" rel="stylesheet">
    <link href="/res/open-iconic/font/css/open-iconic-bootstrap.min.css" rel="stylesheet">
    <!-- JS Scripts -->

    <!-- load jQuery 1.11.0 -->
    <script src="/res/js/jquery_v1.11.0.min.js"></script>
    <script src="/res/js/jquery.websocket-0.0.1.js"></script>
    <script src="/res/js/jquery.json-2.4.min.js"></script>
    <script src="/res/js/page_scripts/ScoringWebsocket.js"></script>
    <script type="text/javascript">
        $jQuery_1_11_0 = $.noConflict(true);
    </script>

    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"
            integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"
            crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js"
            integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4"
            crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js"
            integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1"
            crossorigin="anonymous"></script>



    <!-- Scoring System Script -->
    <script src="/res/js/ftcScoring.js"></script>

    `)
	buffer.WriteString(`
<script src="/res/js/page_scripts/match_play.js"></script>
<script src="/res/js/page_scripts/match_timing.js"></script>
`)

	buffer.WriteString(`

    <title>`)
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
                        </div>
                    </li>
                    <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMatch" data-toggle="dropdown"
                           aria-haspopup="true" aria-expanded="false">
                            Match
                        </a>
                        <div class="dropdown-menu">
                            <a class="dropdown-item" href="/match/play">Play</a>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</div>

<div class="container-fluid">
    `)
	buffer.WriteString(`
<div class="row">
    <div class="col-lg-2">
    </div>
    <div class="col-lg">
        <div class="card bg-light mb-3">
            <div class="card-header">
                Match Status
            </div>
            <div class="card-body">
                matchState: <span id="matchState"></span>
                matchTime: <span id="matchTime"></span>

                <button class="btn btn-outline-primary" id="startMatch" onclick="startMatch();">Start Match</button>
                <button class="btn btn-outline-warning" id="togglePause" onclick="togglePause();">Pause Match</button>
                <button class="btn btn-outline-danger" id="abortMatch" onclick="abortMatch();">AbortMatch</button>
                <button class="btn btn-outline-warning" id="discardResults" onclick="discardResults();">Discard Results</button>
                <button class="btn btn-outline-info" id="commitResults" onclick="commitResults();">Commit Results</button>
            </div>
        </div>
    </div>
    <div class="col-lg-2">
    </div>
</div>

<script>
    $(function () {

    })
</script>

`)

	buffer.WriteString(`
</div>
</body>
</html>`)

}
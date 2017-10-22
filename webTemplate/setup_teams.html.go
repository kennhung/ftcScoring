// Code generated by hero.
// source: C:\work\src\github.com\kennhung\ftcScoring\webTemplate\setup_teams.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"fmt"

	"github.com/kennhung/ftcScoring/model"
	"github.com/shiyanhui/hero"
)

func Setup_Teams(teams []model.Team, buffer *bytes.Buffer) {
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

    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"
            integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"
            crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js"
            integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4"
            crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js"
            integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1"
            crossorigin="anonymous"></script>
    <script src="/res/js/bootstrap-number-input.js"></script>
    <!-- Scoring System Script -->
    <script src="/res/js/ftcScoring.js"></script>

    `)

	buffer.WriteString(`

    <title>`)
	buffer.WriteString(`
Teams
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
	buffer.WriteString(`
<div class="row">
    <div class="col-lg-2">
    </div>
    <div class="col-lg">
        <div class="card ">
            <div class="card-header">
                Teams
            </div>
            <div class="card-body">
                <table class="table table-striped table-hover">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Affiliation</th>
                        <th>City</th>
                        <th>State</th>
                        <th>Country</th>
                    </tr>
                    </thead>
                    <tbody>
                    `)
	for _, team := range teams {

		buffer.WriteString(`
                    <tr>
                        <th scope="row">`)
		hero.EscapeHTML(fmt.Sprint(team.Id), buffer)
		buffer.WriteString(`</th>
                        <td>`)
		hero.EscapeHTML(team.Name, buffer)
		buffer.WriteString(`</td>
                        <td>`)
		hero.EscapeHTML(team.Affiliation, buffer)
		buffer.WriteString(`</td>
                        <td>`)
		hero.EscapeHTML(team.City, buffer)
		buffer.WriteString(`
                        <td>`)
		hero.EscapeHTML(team.State, buffer)
		buffer.WriteString(`</td>
                        <td>`)
		hero.EscapeHTML(team.Country, buffer)
		buffer.WriteString(`</td>
                    </tr>
                    `)
	}
	buffer.WriteString(`
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    <div class="col-lg-2">
    </div>
</div>

`)

	buffer.WriteString(`
</div>
</body>
</html>`)

}
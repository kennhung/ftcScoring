// Code generated by hero.
// source: C:\work\src\github.com\kennhung\ftcScoring\webTemplate\setup_settings.html
// DO NOT EDIT!
package template

import "github.com/shiyanhui/hero"

import (
	"bytes"
	"fmt"

	"github.com/kennhung/ftcScoring/model"
)

func Setup_settings(eventSettings *model.EventSettings, buffer *bytes.Buffer) {
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
    <script src="/res/js/bootstrap-number-input.js"></script>
    <!-- Scoring System Script -->
    <script src="/res/js/ftcScoring.js"></script>

    `)
	buffer.WriteString(`
<script src="/res/datePicker/picker.js"></script>
<script src="/res/datePicker/picker.date.js"></script>
<link href="/res/datePicker/classic.css" rel="stylesheet">
<link href="/res/datePicker/classic.date.css" rel="stylesheet">
`)

	buffer.WriteString(`

    <title>`)
	buffer.WriteString(`
Event Setting
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
        <div class="card bg-light mb-3">
            <div class="card-header">
                Event Settings
            </div>
            <div class="card-body">
                <form id="eventSetting" action="/setup/settings" method="POST" novalidate>
                    <div class="form-group">
                        <label for="name">Event Name</label>
                        <input name="name" type="text" class="form-control" id="name" placeholder="Enter Event Name"
                               value="`)
	hero.EscapeHTML(eventSettings.Name, buffer)
	buffer.WriteString(`">
                    </div>
                    <div class="form-group">
                        <label for="region">Event Region</label>
                        <input name="region" type="text" class="form-control" id="region" placeholder="Region"
                               value="`)
	hero.EscapeHTML(eventSettings.Region, buffer)
	buffer.WriteString(`">
                    </div>
                    <div class="form-group">
                        <label for="type">Event Type</label>
                        <select  name="type" class="form-control" id="type">
                            `)

	var types = []string{"Championship", "Qualifer",
		"Meet", "League Tournament", "Scrimmage", "Other"}

	for i := 0; i < len(types); i++ {
		option := types[i]
		if option == eventSettings.Type {

			buffer.WriteString(`
                            <option selected="selected">
                                `)
		} else {

			buffer.WriteString(`
                            <option>
                                `)
		}
		hero.EscapeHTML(option, buffer)
		buffer.WriteString(`
                            </option>
                            `)

	}

	buffer.WriteString(`
                        </select >
                    </div>
                    <div class="form-group">
                        `)

	var timestr = ""
	year, monthM, day := eventSettings.Date.Date()
	var month int = int(monthM)
	if month < 10 {
		timestr += fmt.Sprint("0", month)
	} else {
		timestr += fmt.Sprint(month)
	}
	timestr += fmt.Sprint("/")
	if day < 10 {
		timestr += fmt.Sprint("0", day)
	} else {
		timestr += fmt.Sprint(day)
	}
	timestr += fmt.Sprint("/", year)

	buffer.WriteString(`
                        <label for="date">Date</label>
                        <div class="input-group">
                            <input name="date" id="date" type="text" class="form-control datepicker" placeholder="mm/dd/yyyy"
                                   value="`)
	hero.EscapeHTML(timestr, buffer)
	buffer.WriteString(`" data-toggle="popover" data-placement="left" data-content="Wrong Date format">
                            <span class="input-group-addon"  id="openPicker"><span class="oi oi-calendar" title="calendar"
                                                                                   aria-hidden="true"></span></span>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="background_color">Background Color</label>
                        <input name="background_color" type="text" class="form-control" id="background_color" placeholder="Background Color"
                               value="`)
	hero.EscapeHTML(eventSettings.DisplayBackgroundColor, buffer)
	buffer.WriteString(`">
                    </div>
                    <div class="form-check">
                        <label class="form-check-label" for="overlay">
                            <input name="overlay" type="checkbox" class="form-check-input" id="overlay"
                                   `)
	if eventSettings.DisplayOverlayMode {
		buffer.WriteString(`
                            checked
                            `)
	}
	buffer.WriteString(` > Video Overlay
                        </label>
                    </div>
                    <button id="send" type="button" class="btn btn-primary">Save</button>
                </form>

            </div>
        </div>
    </div>
    <div class="col-lg-2">
    </div>
</div>

<script>
    $(document).ready(function () {
        $('.alert').hide();
        var $input = $('.datepicker').pickadate({
            // Escape any “rule” characters with an exclamation mark (!).
            format: 'mm/dd/yyyy',
            formatSubmit: 'mm/dd/yyyy',
            editable :true
        })
        var picker = $input.pickadate('picker');

        $("#openPicker").click(function () {
            picker.open(false);
        })

        $("#send").click(function () {
            if (isValidDate($("#date").val())) {
                $("#eventSetting").submit();
            }
            else {
                $("#date").popover('toggle');
                $("#date").addClass('is-invalid');
            }
        })
    })
</script>

`)

	buffer.WriteString(`
</div>
</body>
</html>`)

}

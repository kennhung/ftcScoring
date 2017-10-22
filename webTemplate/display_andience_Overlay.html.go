// Code generated by hero.
// source: C:\work\src\github.com\kennhung\ftcScoring\webTemplate\display_andience_Overlay.html
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/kennhung/ftcScoring/model"
	"github.com/shiyanhui/hero"
)

func Display_Audience_overlay(EventSetting *model.EventSettings, buffer *bytes.Buffer) {
	buffer.WriteString(`

<!DOCTYPE html>
<html>
<head>
    <title>Audience Display - FTC Scoring </title>
    <link rel="stylesheet" href="/res/css/bootstrap.min.css" />

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

    <script src="/res/js/page_scripts/match_timing.js"></script>
    <script src="/res/js/page_scripts/audience_display.js"></script>

</head>
<body style="background-color:`)
	hero.EscapeHTML(EventSetting.DisplayBackgroundColor, buffer)
	buffer.WriteString(`">


<audio id="match-start" src="/res/audio/match_start.wav" preload="auto"></audio>
<audio id="auto-end" src="/res/audio/auto_end.wav" preload="auto"></audio>
<audio id="match-end" src="/res/audio/match_end.wav" preload="auto"></audio>
<audio id="match-abort" src="/res/audio/match_abort.wav" preload="auto"></audio>
<audio id="match-resume" src="/res/audio/match_resume.wav" preload="auto"></audio>
<audio id="match-endgame" src="/res/audio/match_endgame.wav" preload="auto"></audio>
</body>
</html>`)

}

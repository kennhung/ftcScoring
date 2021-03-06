// Code generated by hero.
// source: C:\work\src\github.com\kennhung\ftcScoring\webTemplate\play_listTeams.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"fmt"

	"github.com/kennhung/ftcScoring/model"
	"github.com/shiyanhui/hero"
)

func Match_Play_ListTeams(matchs []model.Match, currentMatch *model.Match, mode int, buffer *bytes.Buffer) {
	buffer.WriteString(`

<table class="table table-sm table-hover">
    <thead>
    <tr>
        <th scope="col">#</th>
        <th scope="col">Time</th>
        <th scope="col">Red</th>
        <th scope="col">Blue</th>
        `)

	if mode == 0 {

		buffer.WriteString(`
        <th scope="col">Load</th>

        `)
	} else if mode == 1 {
		buffer.WriteString(`
        <th scope="col">RedScore</th>
        <th scope="col">BlueScore</th>
        <th scope="col">Action</th>
        `)
	}
	buffer.WriteString(`
    </tr>
    </thead>
    <tbody>
    `)

	for _, match := range matchs {

		var matchName = ""

		switch match.Type {
		case "practice":
			matchName += fmt.Sprint("P-")
		case "qualification":
			matchName += fmt.Sprint("Q-")
			//TODO Playoff
		}
		matchName += fmt.Sprint(match.DisplayName)

		var Red = ""
		var Blue = ""
		var time = ""

		time += fmt.Sprint(match.Time.Hour())
		time += fmt.Sprint(":")
		if match.Time.Minute() < 10 {
			time += fmt.Sprint("0")
			time += fmt.Sprint(match.Time.Minute())
		} else {
			time += fmt.Sprint(match.Time.Minute())
		}

		Red += fmt.Sprint(match.Red1)
		if match.Red1 < 10 {
			Red += fmt.Sprint(" ")
		}
		Red += fmt.Sprint(" ")
		Red += fmt.Sprint(match.Red2)
		if match.Red2 < 10 {
			Red += fmt.Sprint(" ")
		}

		Blue += fmt.Sprint(match.Blue1)
		Blue += fmt.Sprint(" ")
		Blue += fmt.Sprint(match.Blue2)
		var matchID = fmt.Sprint(match.Id)
		var classStr = ""

		if match.Id == currentMatch.Id {
			classStr += fmt.Sprint("table-warning")
		}

		buffer.WriteString(`
    <tr class="`)
		hero.EscapeHTML(classStr, buffer)
		buffer.WriteString(`">
        <th scope="row">`)
		hero.EscapeHTML(matchName, buffer)
		buffer.WriteString(`</th>
        <td>`)
		hero.EscapeHTML(time, buffer)
		buffer.WriteString(`</td>
        <td class="text-danger">`)
		hero.EscapeHTML(Red, buffer)
		buffer.WriteString(`</td>
        <td class="text-primary">`)
		hero.EscapeHTML(Blue, buffer)
		buffer.WriteString(`</td>
        `)

		if mode == 0 {

			buffer.WriteString(`
        <td><a class="btn btn-primary btn-sm btn-load" id="`)
			hero.EscapeHTML(matchID, buffer)
			buffer.WriteString(`" href="/match/play/`)
			hero.EscapeHTML(matchID, buffer)
			buffer.WriteString(`/load">Load</a></td>
        `)
		} else if mode == 1 {

			buffer.WriteString(`
        <td class="text-danger"></td>
        <td class="text-primary"></td>
        <td><a class="btn btn-primary btn-sm btn-load" id="`)
			hero.EscapeHTML(matchID, buffer)
			buffer.WriteString(`" href="/match/review/`)
			hero.EscapeHTML(matchID, buffer)
			buffer.WriteString(`/edit">Edit</a></td>
        `)
		}
		buffer.WriteString(`
    </tr>
    `)

	}

	buffer.WriteString(`
    </tbody>
</table>`)

}

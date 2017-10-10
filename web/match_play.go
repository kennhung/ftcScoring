package web

import (
	"net/http"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"time"
	"log"
)

func (web *Web) matchPlayHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)

	template.Match_Play("test", buffer)
	w.Write(buffer.Bytes())
}

func (web *Web) matchPlayWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	websocket, err := NewWebsocket(w, r)

	if err != nil {
		handleWebErr(w, err)
		return
	}
	defer websocket.Close()
	for {


		err = websocket.Write("test", web.arena.EventSettings)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			return
		}
		time.Sleep(time.Millisecond * 10)

	}
}

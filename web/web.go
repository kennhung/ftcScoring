package web

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kennhung/ftcScoring/arena"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"fmt"
)

type Web struct {
	arena *arena.Arena
}

func NewWeb(arena *arena.Arena) *Web {
	web := &Web{arena: arena}
	return web
}

func (web *Web) ServeWebInterface(webPort int) {
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res/"))))
	http.Handle("/", web.newHandler())

	// Start Server
	log.Printf("Serving HTTP requests on port %d", webPort)
	log.Print(fmt.Sprintf(":%d", webPort))
	http.ListenAndServe(fmt.Sprintf(":%d", webPort), nil)
}

func handleWebErr(w http.ResponseWriter, err error) {
	http.Error(w, "Internal server error: "+err.Error(), 500)
}


func (web *Web)newHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", web.IndexHandler).Methods("GET")
	router.HandleFunc("/setup/settings",web.setupsettingsGetHandler).Methods("GET")
	router.HandleFunc("/setup/settings",web.setupsettingsPOSTHandler).Methods("POST")
	router.HandleFunc("/match/play",web.matchPlayHandler).Methods("GET")
	router.HandleFunc("/match/play/socket",web.matchPlayWebsocketHandler).Methods("GET")
	return router
}

func (web *Web)IndexHandler(w http.ResponseWriter, r *http.Request){
	buffer := new(bytes.Buffer)
	template.Index("Test",buffer)

	w.Write(buffer.Bytes())
}
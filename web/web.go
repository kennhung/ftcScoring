package web

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kennhung/ftcScoring/arena"
	"bytes"
	"github.com/kennhung/ftcScoring/webTemplate"
	"time"
	"fmt"
)

type Web struct {
	arena *arena.Arena
}

func NewWeb(arena *arena.Arena) *Web {
	web := &Web{arena: arena}
	log.Print(web.arena)
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


func (web *Web)newHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", web.IndexHandler).Methods("GET")
	return router
}

func (web *Web)IndexHandler(w http.ResponseWriter, r *http.Request){
	buffer := new(bytes.Buffer)
	template.Display(fmt.Sprint(time.Now()), buffer)

	w.Write(buffer.Bytes())
}
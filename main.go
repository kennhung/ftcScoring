package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
	"github.com/kennhung/ftcScoring/web"
	"github.com/kennhung/ftcScoring/arena"
)

const eventDBPATH = "./foo.db"
const webPort = 80;

func main() {
	log.Print("Scoring System Starting at",time.Now())
	arena, err := arena.NewArena(eventDBPATH)
	if err != nil {
		log.Fatalln("Error during startup: ", err)
	}

	web := web.NewWeb(arena)
	go web.ServeWebInterface(webPort)

	arena.Run()
}
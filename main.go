package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
	"github.com/kennhung/ftcScoring/web"
)

const eventDBPATH = "./foo.db"
const webPort = 8080;

func main() {
	log.Print("Scoring System Starting at",time.Now())

	web := web.NewWeb(nil)
	web.ServeWebInterface(webPort)

}
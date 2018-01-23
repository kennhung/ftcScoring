package web

import (
	"fmt"
	"net"
	"encoding/json"
	"log"
	"io"
	"strings"
	"bytes"
)

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
)

type TCPdata struct {
	messageType string
	message     []byte
}

func (web *Web) ServeSocketInterface(socketPort int) (error) {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+fmt.Sprint(socketPort))
	if err != nil {
		return err
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Printf("Serving TCP Socket requests on port %d", socketPort)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		// Handle connections in a new goroutine.
		go web.handleRequest(conn)
	}
}

// Handles incoming requests.
func (web *Web) handleRequest(conn net.Conn) {

	defer conn.Close()
	log.Print("Commected ", conn.RemoteAddr())

	matchLoadTeamsListener := web.arena.MatchLoadTeamsChannel.Listen()
	defer close(matchLoadTeamsListener)
	matchTimeListener := web.arena.MatchTimeChannel.Listen()
	defer close(matchTimeListener)
	reloadDisplaysListener := web.arena.ReloadDisplaysChannel.Listen()
	defer close(reloadDisplaysListener)

	go func() {
		for {
			data := TCPdata{}

			select {
			case _, ok := <-matchLoadTeamsListener:
				if !ok {
					return
				}
				data.messageType = "reload"
				data.message = nil
			case matchTimeSec, ok := <-matchTimeListener:
				if !ok {
					return
				}
				data.messageType = "matchTime"
				data.message, _ = json.Marshal(MatchTimeMessage{web.arena.MatchState, matchTimeSec.(int)})
			case _, ok := <-reloadDisplaysListener:
				if !ok {
					return
				}
				data.messageType = "reload"
				data.message = nil
			}

			_, err := conn.Write([]byte(encodeSendingString(data)))
			if err != nil {
				// The client has probably closed the connection; nothing to do here.
				return
			}
		}
	}()

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				// Client has closed the connection; nothing to do here.
				return
			}
			log.Printf("TCPSocket error: %s", err)
			return
		}

		messageType, data := decodeSendingString(string(buf))
		if data == nil {
			return
		}
		switch messageType {
		case "ping":

			log.Print(data)
		default:
			log.Panic(fmt.Sprintf("Invalid message type '%s'.", messageType))
		}
	}
}

func encodeSendingString(data TCPdata) (string) {
	sendingString := ""
	sendingString += fmt.Sprint(data.messageType)
	sendingString += fmt.Sprint(";")
	sendingString += fmt.Sprint(data.message)
	return sendingString
}

func decodeSendingString(dataString string) (string, map[string]interface{}) {
	s := strings.Split(dataString, ";")
	u := map[string]interface{}{}
	err := json.Unmarshal(bytes.Trim([]byte(s[1]), "\x00"), &u)
	if err != nil {
		log.Panic("Json unmarshal error : ", err)
		return s[0], nil
	}
	return s[0], u
}

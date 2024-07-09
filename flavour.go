package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/r3labs/sse/v2"
)

var curFlavour = 0
var flavoursUrl = [3]string{"https://www.nyan.cat/cats/original.gif", "https://www.nyan.cat/cats/gb.gif", "https://www.nyan.cat/cats/jazz.gif"}

var mu sync.Mutex

func ChangeFlavoursTicker(server *sse.Server) {
	for {
		mu.Lock()
		curFlavour = (curFlavour + 1) % len(flavoursUrl)
		mu.Unlock()

		server.Publish("flavour", &sse.Event{
			Data: []byte(flavoursUrl[curFlavour]),
		})

		time.Sleep(10 * time.Second)
	}
}

func SetCorsHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func ChangeFlavoursNow(server *sse.Server) {

	mu.Lock()
	curFlavour = (curFlavour + 1) % len(flavoursUrl)
	mu.Unlock()

	log.Println("Change flavour manually to " + flavoursUrl[curFlavour])
	server.Publish("flavour", &sse.Event{
		Data: []byte(flavoursUrl[curFlavour]),
	})

}

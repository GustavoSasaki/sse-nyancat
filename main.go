package main

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

func main() {
	logFile := startLogger()
	defer logFile.Close()

	server := sse.New()
	server.Headers["Access-Control-Allow-Origin"] = "*"
	server.CreateStream("flavour")

	router := http.NewServeMux()
	router.HandleFunc("/events", server.ServeHTTP)

	go ChangeFlavoursTicker(server)

	router.HandleFunc("/change-flavour", func(w http.ResponseWriter, r *http.Request) {
		SetCorsHeader(&w)

		// handle OPTION request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ChangeFlavoursNow(server)
	})

	http.ListenAndServe(":8080", router)
}

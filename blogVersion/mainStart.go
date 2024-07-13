package blogVersion

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

// <start>
// <show>
func main() {

	server := sse.New()
	server.AutoReplay = false
	server.Headers["Access-Control-Allow-Origin"] = "*"
	server.CreateStream("flavour")

	router := http.NewServeMux()
	router.HandleFunc("/events", server.ServeHTTP)

	http.ListenAndServe(":8080", router)
}

// </show>

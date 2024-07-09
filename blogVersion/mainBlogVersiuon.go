package blogVersion

import (
	"net/http"
	"time"

	"github.com/r3labs/sse/v2"
)

// <start>
func main() {

	server := sse.New()
	server.Headers["Access-Control-Allow-Origin"] = "*"
	server.CreateStream("flavour")

	router := http.NewServeMux()
	router.HandleFunc("/events", server.ServeHTTP)

	http.ListenAndServe(":8080", router)
	// <sse>
	var curFlavour = 0
	var flavoursUrl = [3]string{"https://www.nyan.cat/cats/original.gif", "https://www.nyan.cat/cats/gb.gif", "https://www.nyan.cat/cats/jazz.gif"}
	go func() {
		for {
			curFlavour = (curFlavour + 1) % len(flavoursUrl)

			server.Publish("flavour", &sse.Event{
				Data: []byte(flavoursUrl[curFlavour]),
			})

			time.Sleep(10 * time.Second)
		}
	}()
	// </sse>
}

// </start>

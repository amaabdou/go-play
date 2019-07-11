package server

import (
	"fmt"
	"log"
	"net/http"
)

func Listen(port int64) {
	http.HandleFunc("/", requestResponseHandler)
	log.Printf("Listening on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

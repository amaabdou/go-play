package server

import (
	"fmt"
	"github.com/amaabdou/go-play-filesystem-http-browser/browser"
	"net/http"
)

func RequestResponseHandler(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	requestedPath := httpRequest.URL.Path
	fmt.Fprintf(httpResponseWriter, browser.Browser(requestedPath))
}

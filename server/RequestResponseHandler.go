package server

import (
	"fmt"
	"github.com/amaabdou/go-play-filesystem-http-browser/browser"
	"net/http"
)

func RequestResponseHandler(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	requestedPath := httpRequest.URL.Path
	browserPage := browser.Browser(requestedPath)
	fmt.Fprintf(httpResponseWriter, browser.PageRenderer(browserPage))
}

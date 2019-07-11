package browser

import "fmt"

func pageRenderer( page page)  string {
	return fmt.Sprintf("<html><title>%s</title><bod>%s</body></html>", page.Title, page.Body)
}
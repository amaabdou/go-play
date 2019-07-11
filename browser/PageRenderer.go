package browser

import "fmt"

func PageRenderer( page Page)  string {
	return fmt.Sprintf("<html><title>%s</title><bod>%s</body></html>", page.Title, page.Body)
}
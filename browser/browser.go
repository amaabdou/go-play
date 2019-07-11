package browser

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Browser(path string) string {
	page := allBrowser(path)
	return pageRenderer(page)
}

func allBrowser(path string) page {
	var page page

	fi, err := os.Stat(path)
	if err != nil {
		page.Body = err.Error()
		return page
	}

	if fi.Mode().IsDir() {
		return dirBrowser(path)
	}

	if fi.Mode().IsRegular() {
		return fileBrowser(path)
	}

	page.Body = fmt.Sprintf("un supported type <br>%s %s", path, fi.Mode().String())
	return page
}

func dirBrowser(path string) page {
	var page page

	dirContents, err := ioutil.ReadDir(path)
	if err != nil {
		page.Title = err.Error()
		return page
	}

	page.Title = path

	if path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}

	dirList := fmt.Sprintf(" Found %d items", len(dirContents))
	dirList += "<ul>"
	dirList += fmt.Sprintf("<li> <a href='%s/../'>..</a>", path)
	for _, dirContent := range dirContents {
		fullPAth := fmt.Sprint(path, "/", dirContent.Name())
		dirList += fmt.Sprintf("<li>%s <a href='%s'>%s</a>", dirContent.Mode().String(), fullPAth, fullPAth)
	}
	dirList += "</ul>"

	page.Body = fmt.Sprintf("<pre>%s</pre>", dirList)

	return page
}

func fileBrowser(path string) page {
	var page page

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		page.Title = err.Error()
		return page
	}

	page.Title = path
	page.Body = fmt.Sprintf("<pre>%s</pre>", string(fileContent))

	return page
}
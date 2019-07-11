package browser

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Browser(path string) Page {
	var x Page

	fi, err := os.Stat(path)
	if err != nil {
		x.Body = err.Error()
		return x
	}

	if fi.Mode().IsDir() {
		return DirectoryBrowser(path)
	}

	if fi.Mode().IsRegular() {
		return FileBrowser(path)
	}

	x.Body = fmt.Sprintf("un supported type <br>%s %s", path, fi.Mode().String())
	return x
}

func DirectoryBrowser(path string) Page {
	var x Page

	dirContents, err := ioutil.ReadDir(path)
	if err != nil {
		x.Title = err.Error()
		return x
	}

	x.Title = path

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

	x.Body = fmt.Sprintf("<pre>%s</pre>", dirList)

	return x
}

func FileBrowser(path string) Page {
	var x Page

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		x.Title = err.Error()
		return x
	}

	x.Title = path
	x.Body = fmt.Sprintf("<pre>%s</pre>", string(fileContent))

	return x
}
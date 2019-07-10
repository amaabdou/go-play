package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path
	p := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1> <bod>%s</body>", p.Title, p.Body)}

type Page struct {
	Title string
	Body string
}

func loadPage(s string) Page {
	var x Page

	fi, err := os.Stat(s)
	if err !=nil {
		x.Body = err.Error()
		return  x
	}

	if fi.Mode().IsDir() {
		return getDirContent(s)
	}

	if fi.Mode().IsRegular() {
		return getFileContent(s)
	}

	x.Body = fmt.Sprintf("un supported type <br>%s %s", s, fi.Mode().String())

	return  x
}

func getDirContent(path string) Page {
	var x Page

	dirContents, err := ioutil.ReadDir(path)
	if err !=nil {
		x.Title = err.Error()
		return  x
	}

	x.Title = path

	if path[len(path)-1] == '/' {
		path = path[0:len(path)-1]
	}

	dirList := fmt.Sprintf(" Found %d items", len(dirContents))
	dirList += "<ul>"
	dirList += fmt.Sprintf("<li> <a href='%s/../'>..</a>", path)
	for _,dirContent := range dirContents {
		fullPAth := fmt.Sprint(path,"/",dirContent.Name())
		dirList += fmt.Sprintf("<li>%s <a href='%s'>%s</a>", dirContent.Mode().String(), fullPAth, fullPAth)
	}
	dirList += "</ul>"

	x.Body = fmt.Sprintf("<pre>%s</pre>" , dirList)

	return  x
}


func getFileContent(path string) Page {
	var x Page

	fileContent, err := ioutil.ReadFile(path)
	if err !=nil {
		x.Title = err.Error()
		return  x
	}

	x.Title = path
	x.Body = fmt.Sprintf("<pre>%s</pre>" , string(fileContent))

	return  x
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE: go run test.go PORT")
	}

	port, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err !=nil {
		log.Fatal("USAGE: go run test.go PORT")
	}

	http.HandleFunc("/", handler)
	log.Printf("Listening on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

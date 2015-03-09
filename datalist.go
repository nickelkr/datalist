package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"html/template"
	"os"
)

type Entry struct {
	Name 				string
	Link 				string
	Description string
}

var logger = setupLog()
var templates = template.Must(template.ParseFiles("views/links.html"))
var testEntries = []Entry{{"This", "one", "is a test"}}

// SetupLog initializes our Logger and returns a pointer to the Logger
func setupLog() *log.Logger {
	logFile, err := os.Create("./datalist.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log := log.New(logFile, "datalist", log.Lshortfile)
	return log
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

// Render executes the template for a given name
func render(w http.ResponseWriter, tmpl string, entries []Entry) {
	err := templates.ExecuteTemplate(w, tmpl+".html", entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Index returns the front-page of newly registered data sources
func index(w http.ResponseWriter, r *http.Request) {
  render(w, "links", testEntries)
}

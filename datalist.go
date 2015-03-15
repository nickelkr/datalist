package main

import (
	"github.com/gorilla/mux"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"html/template"
	"os"
)

type Source struct {
	Id					bson.ObjectId `bson:"_id"`
	Name 				string 				`bson:"name"`
	Link 				string				`bson:"link"`
	Description string				`bson:"description"`
}

var logger = setupLog()
var templates = template.Must(template.ParseFiles("views/links.html",
																									"views/new.html"))
var testSources = []Source{{"This", "one", "is a test"}}

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
	r.HandleFunc("/new/", createHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

// Render executes the template for a given name
func render(w http.ResponseWriter, tmpl string, sources []Source) {
	err := templates.ExecuteTemplate(w, tmpl+".html", sources)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateHandler returns a blank form for a new source
func createHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "new", nil)
}

// Index returns the front-page of newly registered data sources
func index(w http.ResponseWriter, r *http.Request) {
  render(w, "links", testEntries)
}

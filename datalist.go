package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var logger = setupLog()

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

func index(w http.ResponseWriter, r *http.Request) {
  logger.Printf("In index")
}

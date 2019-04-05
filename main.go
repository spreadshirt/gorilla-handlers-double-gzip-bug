package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	gorilla := mux.NewRouter()

	gorilla.Handle("/", promhttp.Handler())
	gorilla.Use(handlers.CompressHandler)

	fmt.Println("gorilla/handlers.CompressHandler: Listening on :9999")
	go func() {
		err := http.ListenAndServe(":9999", gorilla)
		if err != nil {
			log.Fatal(err)
		}
	}()

	nytimes := mux.NewRouter()

	nytimes.Handle("/", promhttp.Handler())
	nytimes.Use(gziphandler.GzipHandler)

	fmt.Println("nytimes/gziphandler.GzipHandler: Listening on :19999")
	err := http.ListenAndServe(":19999", nytimes)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", promhttp.Handler())
	r.Use(handlers.CompressHandler)

	fmt.Println("Listening on :9999")
	http.ListenAndServe(":9999", r)
}

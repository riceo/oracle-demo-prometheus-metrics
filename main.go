package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requests = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Requests to the demo endpoint",
		})
)

func handler(w http.ResponseWriter, r *http.Request) {
	requests.Inc()
	w.Write([]byte("This service is _so_ micro."))
}

func main() {
	fmt.Printf("HTTP server is running on port 8080!\n")
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

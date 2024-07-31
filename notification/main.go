package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	http.HandleFunc("/notification", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Notification Service"))
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Notification service is running on :8084")
	http.ListenAndServe(":8084", nil)
}

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
	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Payment Service"))
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Payment service is running on :8082")
	http.ListenAndServe(":8082", nil)
}

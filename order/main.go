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
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Order Service"))
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Order service is running on :8083")
	http.ListenAndServe(":8083", nil)
}

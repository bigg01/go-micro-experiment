package main

import (
	"log"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go/config"
)

var authCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "auth_requests_total",
		Help: "Total number of /auth requests",
	},
)

func init() {
	// Register the counter with Prometheus's default registry.
	prometheus.MustRegister(authCounter)
	log.SetOutput(os.Stdout)
}

func main() {
	// Initialize Jaeger Tracer
	cfg := config.Configuration{
		ServiceName: "auth-service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "0.0.0.0:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// Start a new span
		span := tracer.StartSpan("auth-endpoint")
		defer span.Finish()

		authCounter.Inc() // Increment the counter
		w.Write([]byte("Auth Service"))
	})

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Auth service is running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Microservices Endpoints</title>
            </head>
            <body>
                <h1>Microservices Endpoints</h1>
                <ul>
					<li><a href="http://localhost:8081/auth">Auth Service</a></li>
					<li><a href="http://localhost:8081/metrics">Auth Service Metrics</a></li>
					<li><a href="http://localhost:8082/payment">Payment Service</a></li>
                    <li><a href="http://localhost:8083/order">Order Service</a></li>
                    <li><a href="http://localhost:8084/notification">Notification Service</a></li>

                </ul>
            </body>
            </html>
        `)
	})

	http.ListenAndServe(":8080", nil)
}

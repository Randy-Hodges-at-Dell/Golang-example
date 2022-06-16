package main

import (
        "net/http"
        "github.com/prometheus/client_golang/prometheus/promhttp"
		"fmt"
)

func main() {
	fmt.Println("hello world")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
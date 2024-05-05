package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		pMux := http.NewServeMux()
		promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
		pMux.Handle("/metrics", promHandler)

		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	Init()
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cavaliercoder/rpi_export/pkg/export/prometheus"
)

var (
	addr = flag.String("addr", ":9110", "Address to listen on")
)

func main() {
	flag.Parse()

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		if err := prometheus.Write(w); err != nil {
			http.Error(w, fmt.Sprintf("Error generating metrics: %v", err), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<html>
<head><title>Raspberry Pi Exporter</title></head>
<body>
<h1>Raspberry Pi Exporter</h1>
<p><a href="/metrics">Metrics</a></p>
</body>
</html>`)
	})

	log.Printf("Starting rpi_exporter on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
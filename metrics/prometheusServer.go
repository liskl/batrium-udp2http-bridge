package metrics

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// starts prometheuse server
func StartPrometheusServer(addr string) {
	if addr == "" {
		addr = ":9001"
	}
	r := mux.NewRouter()
	r.Path("/metrics").Handler(promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(addr, r); err != nil {
			panic(err)
		}
	}()
}
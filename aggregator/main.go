package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/bzawada1/location-app-obu-service/types"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	store := NewMemoryStore()
	svc := NewInvoiceAggregator(store)
	svc = NewLogMiddleware(svc)
	svc = NewMetricsMiddleware(svc)
	grpcAddr := os.Getenv("AGG_GRPC_ENDPOINT")
	httpAddr := os.Getenv("AGG_HTTP_ENDPOINT")
	go makeGRPCTransport(grpcAddr, svc)
	makeHTTPTransport(httpAddr, svc)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	ln, err := net.Listen("TCP", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	server := grpc.NewServer([]grpc.ServerOption{}...)
	types.RegisterAggregatorServer(server, NewAggregatorGRPCServer(svc).server)
	return server.Serve(ln)

}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.HandleFunc("/invoice", handleGetInvoice(svc))
	http.HandleFunc("/invoice/all", handleGetAllInvoice(svc))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listenAddr, nil)
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not supported"})
		}
		values, ok := r.URL.Query()["obu"]
		if !ok {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})
			return
		}
		obuID, err := strconv.Atoi(values[0])
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid OBU ID"})
			return
		}
		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, invoice)
		return
	}
}

func handleGetAllInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, nil)
		return
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not supported"})
		}
		distance := types.Distance{}

		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

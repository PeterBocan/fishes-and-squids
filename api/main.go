package main

// Run this go server on port 8080
// Connect to the Postgres database using the same credentials as in docker-compose.yml and on port
// 5432

import (
	"log"
	"net/http"
	"oo-api/db"
	"oo-api/service"
	"os"
)

// This does not belong to be here. Preferably some other location would be better.
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var (
		DB_HOSTNAME = os.Getenv("DB_HOSTNAME")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	// this would be set in the env.
	if DB_HOSTNAME == "" {
		DB_HOSTNAME = "localhost"
	}
	if DB_PORT == "" {
		DB_PORT = "5432"
	}
	if DB_PASSWORD == "" {
		DB_PASSWORD = "oo"
	}
	if DB_NAME == "" {
		DB_NAME = "oo"
	}
	if DB_USER == "" {
		DB_USER = "oo"
	}

	dbConn, err := db.NewConnection(DB_HOSTNAME, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	if err != nil {
		log.Fatalf("could not establish the connection: %s", err)
	}

	// NOTE: This is very much ad-hoc. Works for very simple modelling. Does only add and change.
	// It does not remove any data, leaving lots of mess behind.
	db.RunMigrations(dbConn)

	// NOTE: not the most perf thing, but works.
	// This makes the startup a bit slower and also a bit problematic but it would be better to have
	// it separately out of the startup of the service. OOS.
	db.RunSeeds(dbConn)

	httpServer := service.New(db.NewDB(dbConn))
	mux := http.NewServeMux()
	mux.HandleFunc("/products", httpServer.GetProducts)

	http.ListenAndServe("localhost:8080", enableCORS(mux))
}

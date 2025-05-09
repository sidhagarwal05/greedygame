package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"greedy-game/internal/db/sqlc"
	"greedy-game/internal/handler"
)

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/test?sslmode=disable&search_path=app"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries := sqlc.New(db)
	h := handler.NewHandler(queries)

	r := mux.NewRouter()
	r.HandleFunc("/v1/delivery", h.Delivery).Methods("GET")

	log.Println("server started at :8080")
	http.ListenAndServe(":8080", r)
}

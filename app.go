package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App basic application structure
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize method
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s passowrd=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run mathod
func (a *App) Run(addr string) {}

package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App basic application structure
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize method
func (a *App) Initialize(user, password, dbname string) {}

// Run mathod
func (a *App) Run(addr string) {}

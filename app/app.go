package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/ypack/rest-api-server/logger"
	"github.com/ypack/rest-api-server/service"
	"log"
	"net/http"
)

// App initialize and run the main application
type App struct {
	Router   *mux.Router
	database *sql.DB
}

// Initialize init the application with routes and database.
// First it tries to connect to the database, and if an error occurs,
// the program exits and show the error in stdout. When the database
// connection is success, then we create a sub-router for the v1 routes.
func (app *App) Initialize(config DatabaseConfig) {
	// Create database connection
	var err error
	app.database, err = sql.Open("mysql", config.ToString())
	if err != nil {
		log.Fatal(err)
	}
	// Add logging middleware
	app.Router.Use(logger.Logger)

	v1group := app.Router.PathPrefix("/v1").Subrouter()
	v1group.HandleFunc("/packages", service.GetPackagesHandler).Methods("GET")
	v1group.HandleFunc("/package", service.GetPackageHandler).Methods("GET")
	v1group.HandleFunc("/package/latest", service.GetLatestPackageHandler).Methods("GET")
}

func (app *App) Run(config ServerConfig) {
	log.Fatal(http.ListenAndServe(config.ToString(), app.Router))
}

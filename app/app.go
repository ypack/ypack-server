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
	router   *mux.Router
	database *sql.DB

	// Version is the current application version
	Version string
}

// Initialize init the application with routes and database.
// First it tries to connect to the database, and if an error occurs,
// the program exits and show the error in stdout. When the database
// connection is success, then we create a sub-router for the v1 routes.
func (app *App) Initialize(config DatabaseConfig) {
	var err error
	app.database, err = sql.Open("mysql", config.ToString())
	if err != nil {
		log.Fatal(err)
	}
	// Create the package service
	ds := service.PackageService{DB: app.database}

	// Create router and add logging middleware
	app.router = mux.NewRouter()
	app.router.Use(logger.Logger)

	v1group := app.router.PathPrefix("/v1").Subrouter()
	v1group.HandleFunc("/packages", ds.GetPackagesHandler).Methods("GET")
	v1group.HandleFunc("/package", ds.GetPackageHandler).Methods("GET")
	v1group.HandleFunc("/package/latest", ds.GetLatestPackageHandler).Methods("GET")
}

// Run starts the Rest API with the given config
func (app *App) Run(config ServerConfig) {
	log.Printf("Server version %s running on %s\n", app.Version, config.ToString())
	log.Fatal(http.ListenAndServe(config.ToString(), app.router))
}

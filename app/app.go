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
	Database *sql.DB
}

// Initialize the application with routes and database
func (app *App) Initialize(config DatabaseConfig) {
	app.Router.Use(logger.Logger)

	v1group := app.Router.PathPrefix("/v1").Subrouter()
	v1group.HandleFunc("/packages", service.GetPackagesHandler).Methods("GET")
	v1group.HandleFunc("/package", service.GetPackageHandler).Methods("GET")
	v1group.HandleFunc("/package/latest", service.GetLatestPackageHandler).Methods("GET")
}

func (app *App) Run(config ServerConfig) {
	log.Fatal(http.ListenAndServe(config.ToString(), app.Router))
}

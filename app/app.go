package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ypack/rest-api-server/logger"
	"github.com/ypack/rest-api-server/service"
	"log"
	"net/http"
	"time"
)

// App initialize and run the main application
type App struct {
	server   *http.Server
	router   *mux.Router
	database *gorm.DB

	// Version is the current application version
	Version string
	// Debug flag to show queries
	Debug bool
}

// Initialize init the application with routes and database.
// First it tries to connect to the database, and if an error occurs,
// the program exits and show the error in stdout. When the database
// connection is success, then we create a sub-router for the v1 routes.
func (app *App) Initialize(dbConfig DatabaseConfig, serverConfig ServerConfig) {
	var err error
	app.database, err = gorm.Open("mysql", dbConfig.ToString())
	if err != nil {
		log.Fatal(err)
	}
	// Disable table name pluralization that uses gorm
	app.database.SingularTable(true)
	if app.Debug {
		app.database.LogMode(true)
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

	// Create server and set timeouts to avoid Slowloris attacks
	app.server = &http.Server{
		Addr:         serverConfig.ToString(),
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      app.router,
	}
}

// Run starts the Rest API with the given config
func (app *App) Run() {
	err := app.database.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server version %s running on %s\n", app.Version, app.server.Addr)
	if err := app.server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

// Shutdown closes the database connection
func (app *App) Shutdown() {
	log.Println("Releasing server resources...")

	// shutdown server
	err := app.server.Close()
	if err != nil {
		log.Fatal(err)
	}

	// close database connection
	err = app.database.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database: Connection closed")
}

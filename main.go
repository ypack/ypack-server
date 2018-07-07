package main

import (
	"github.com/gorilla/mux"
	"github.com/ypack/rest-api-server/app"
)

func main() {
	// Create database config. This will be used to connect to the database
	// with the given user, password and database name
	dbConfig := app.DatabaseConfig{}

	application := app.App{
		Router:   mux.NewRouter(),
		Database: nil,
	}
	// Initialize the application. Connection to the database will occur here.
	// Also, we initialize here the api routes
	application.Initialize(dbConfig)

	// Setup server config and run it
	serverConfig := app.ServerConfig{
		Address: "localhost",
		Port:    8080,
	}
	application.Run(serverConfig)
}

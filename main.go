package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ypack/rest-api-server/app"
	"os"
)

var exitWithError = -1

func main() {
	dbName := flag.String("db-name", "ypack", "Name of the database to connect")
	dbUser := flag.String("db-user", "", "User to connect to the database")
	dbPass := flag.String("db-pass", "", "Password for the given user")

	host := flag.String("host", "localhost", "Host to attach the server")
	port := flag.Int("port", 8080, "Port to listen for requests")

	flag.Parse()

	// Now we need to validate if the required params are set or not.
	// These params are database user and database password because,
	// for the other options, defaults are right
	if *dbUser == "" || *dbPass == "" {
		flag.PrintDefaults()
		os.Exit(exitWithError)
	}

	// Create database config. This will be used to connect to the database
	// with the given user, password and database name
	dbConfig := app.DatabaseConfig{
		User:         *dbUser,
		Password:     *dbPass,
		DatabaseName: *dbName,
	}
	application := app.App{
		Router:   mux.NewRouter(),
		Database: nil,
	}
	// Initialize the application. Connection to the database will occur here.
	// Also, we initialize here the api routes
	application.Initialize(dbConfig)

	// Setup server config and run it
	serverConfig := app.ServerConfig{
		Address: *host,
		Port:    *port,
	}
	application.Run(serverConfig)
}

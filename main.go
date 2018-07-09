package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ypack/rest-api-server/app"
	"os"
)

const (
	exitWithError = -1
	version       = "0.1.0"
	banner        = `

██╗   ██╗██████╗  █████╗  ██████╗██╗  ██╗    ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
╚██╗ ██╔╝██╔══██╗██╔══██╗██╔════╝██║ ██╔╝    ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
 ╚████╔╝ ██████╔╝███████║██║     █████╔╝     ███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
  ╚██╔╝  ██╔═══╝ ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
   ██║   ██║     ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
   ╚═╝   ╚═╝     ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝`
)

func main() {
	dbHost := flag.String("db-host", "localhost", "Host of the database")
	dbPort := flag.Int("db-port", 3306, "Database port to connect")
	dbName := flag.String("db-name", "ypack", "Name of the database to connect")
	dbUser := flag.String("db-user", "", "User to connect to the database")
	dbPass := flag.String("db-pass", "", "Password for the given user")

	serverHost := flag.String("host", "localhost", "Host to attach the server")
	serverPort := flag.Int("port", 8080, "Port to listen for requests")

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
		Host:         *dbHost,
		Port:         *dbPort,
	}
	application := app.App{Version: version, Debug: true}
	// Initialize the application. Connection to the database will occur here.
	// Also, we initialize here the api routes
	application.Initialize(dbConfig)

	// Setup server config and run it
	serverConfig := app.ServerConfig{
		Address: *serverHost,
		Port:    *serverPort,
	}
	fmt.Println(banner)
	application.Run(serverConfig)
}

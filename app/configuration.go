package app

import "fmt"

// DatabaseConfig to connect to the db
type DatabaseConfig struct {
	User         string
	Password     string
	DatabaseName string
	Host         string
	Port         int
}

// ToString transforms the DatabaseConfig struct to a readable format by the application.
// It creates the url needed by the database driver.
func (dc *DatabaseConfig) ToString() string {
	return fmt.Sprintf("jdbc:mariadb://%s:%d/%s?user=%s&password=%s",
		dc.Host, dc.Port, dc.DatabaseName, dc.User, dc.Password)
}

// ServerConfig to start the rest server
type ServerConfig struct {
	Address string
	Port    int
}

// ToString transforms the ServerConfig struct to a readable format by the application
func (sc *ServerConfig) ToString() string {
	return fmt.Sprintf("%s:%d", sc.Address, sc.Port)
}

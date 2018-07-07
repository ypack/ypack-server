package app

import "fmt"

// DatabaseConfig to connect to the db
type DatabaseConfig struct {
	User         string
	Password     string
	DatabaseName string
}

// ServerConfig to start the rest server
type ServerConfig struct {
	Address string
	Port    int
}

func (sc *ServerConfig) ToString() string {
	return fmt.Sprintf("%s:%d", sc.Address, sc.Port)
}

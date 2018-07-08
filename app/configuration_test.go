package app

import "testing"

func TestServerConfig_ToString(t *testing.T) {
	expected := "localhost:8080"
	config := ServerConfig{"localhost", 8080}
	if expected != config.ToString() {
		t.Error("Expected", expected, "but got", config.ToString())
	}
}

func TestDatabaseConfig_ToString(t *testing.T) {
	expected := "root:myPassword@tcp(localhost:3306)/DB"
	config := DatabaseConfig{"root", "myPassword", "DB", "localhost", 3306}
	if expected != config.ToString() {
		t.Error("Expected", expected, "but got", config.ToString())
	}
}

package validate

import "testing"

func TestValidateOperatingSystem(t *testing.T) {
	err1 := IsValidOperatingSystem("hello")
	if err1 == nil {
		t.Error("Expected error but got nil")
	}
	err2 := IsValidOperatingSystem("ubuntu")
	if err2 != nil {
		t.Error("Expected nil error but got", err2)
	}
	err3 := IsValidOperatingSystem("")
	if err3 != nil {
		t.Error("Expected empty but got", err3)
	}
}

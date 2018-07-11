package validate

import "testing"

func TestValidateOperatingSystem(t *testing.T) {
	err1 := IsValidOperatingSystem("hello")
	if err1 == nil {
		t.Error("Expected error but got nil")
	}
	err2 := IsValidOperatingSystem("ubuntu")
	if err2 != nil {
		t.Error("Expected nil error but got:", err2)
	}
	err3 := IsValidOperatingSystem("")
	if err3 == nil {
		t.Error("Expected empty but got:", err3)
	}
}

func TestIsValidPackageName(t *testing.T) {
	err1 := IsValidPackageName("")
	if err1 == nil {
		t.Error("Expected error (empty string) but got:", err1)
	}
	err2 := IsValidPackageName("123456789012345678901234567890")
	if err2 == nil {
		t.Error("Expected error (text lenght > 25) but got:", err2)
	}
	err3 := IsValidPackageName("golang")
	if err3 != nil {
		t.Error("Expected nil but got:", err3)
	}
}

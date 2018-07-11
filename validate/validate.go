package validate

import (
	"fmt"
)

var validOperatingSystems = []string{
	"win",
	"mac",
	"ubuntu",
}

// IsValidOperatingSystem validate if the given OS is know or not. If the
// OS is not valid this method returns an error, otherwise returns nil
func IsValidOperatingSystem(os string) error {
	if os == "" {
		return fmt.Errorf("operating system can't be empty")
	}
	for i := 0; i < len(validOperatingSystems); i++ {
		if validOperatingSystems[i] == os {
			return nil
		}
	}
	return fmt.Errorf("operating system not supported")
}

// IsValidPackageName checks if the given name is valid or not
func IsValidPackageName(name string) error {
	if name == "" {
		return fmt.Errorf("package name can't be empty")
	}
	if len(name) > 25 {
		return fmt.Errorf("invalid package lenght (> 25)")
	}
	return nil
}

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
		return nil
	}
	for i := 0; i < len(validOperatingSystems); i++ {
		if validOperatingSystems[i] == os {
			return nil
		}
	}
	return fmt.Errorf("operating system not supported")
}

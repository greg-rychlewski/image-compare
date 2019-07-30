package flagutil

import "errors"

// Validate user input to command-line flags
func ValidateInputPath(inputPath string) error {
	if inputPath == "" {
		return errors.New("location of input csv file is missing. please specify using the -in flag")
	}

	return nil
}

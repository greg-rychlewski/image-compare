package flagutil

import "errors"

// Validate user input to command-line flags
func ValidateInputPath(inputPath string) error {
	if inputPath == "" {
		return errors.New("Location of input csv file is missing. Please specify using the -in flag.")
	}

	return nil
}

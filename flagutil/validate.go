package flagutil

// Custom error type so that flag errors can cause the program to print flag defaults
type FlagError struct {
	s string
}

func (e *FlagError) Error() string {
	return e.s
}

// Validate user input to command-line flags

func ValidateInputPath(inputPath string) error {
	if inputPath == "" {
		return &FlagError{"Input path is missing"}
	}

	return nil
}

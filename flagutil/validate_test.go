package flagutil

import "testing"

func TestValidateInputPath(t *testing.T) {
	// Test #1: Empty path should produce an error
	err := ValidateInputPath("")

	if err == nil {
		t.Error("Empty input path did not produce an error.")
	}

	// Test #2: Non-empty path should not produce an error
	err = ValidateInputPath("test123")

	if err != nil {
		t.Error("Non-empty input path produces an error.")
	}
}

package flagutil

import "testing"

func TestValidateInputPath(t *testing.T) {
	err := ValidateInputPath("")

	if err == nil {
		t.Error("Empty input path did not produce error.")
	}

	err = ValidateInputPath("test123")

	if err != nil {
		t.Error("Non-empty input path produced error.")
	}
}

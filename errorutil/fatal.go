package errorutil

import (
	"fmt"
	"os"
)

// Exit from program when fatal error occurs

func Exit(err error) {
	fmt.Fprintln(os.Stderr, "(ERROR)", err)
	os.Exit(1)
}

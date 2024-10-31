package log

import (
	"fmt"
	"os"
)

func ExitWithError(msg string) {
	fmt.Fprintf(os.Stderr, "error: %v\n", msg)
	os.Exit(1)
}

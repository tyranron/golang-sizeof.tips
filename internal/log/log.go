package log

import (
	"fmt"
	"os"
)

// Performs printf() of given pattern with given arguments
// to OS standard error output stream (stderr).
func StdErr(pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern, args...)
}

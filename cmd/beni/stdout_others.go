// +build !windows

package main

import (
	"io"
	"os"
)

func getStdoutWriter() io.Writer {
	return os.Stdout
}

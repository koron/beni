package main

import (
	"io"

	"github.com/mattn/go-colorable"
)

func getStdoutWriter() io.Writer {
	return colorable.NewColorableStdout()
}

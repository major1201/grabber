package logging

import (
	"io"
	"os"
	"fmt"
)

type OutputWriter struct {
	Writer io.Writer
	Level int
}

var Writers []OutputWriter

func AddStdout(level int) {
	Writers = append(Writers, OutputWriter{os.Stdout, level})
}

func AddFileWriter(level int, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("log file open error")
	}
	Writers = append(Writers, OutputWriter{f, level})
}

package ax

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// A Null Logger is a logger that doesn't actually log.
type NullLogger struct{}

func (l *NullLogger) Printf(format string, values ...interface{}) {
}

func (l *NullLogger) SetLogger(logger Logger) {
}

func (l *NullLogger) WrappedLogger() Logger {
	return l
}

// StdoutLogger is a FileLogger that logs to stdout
var StdoutLogger = FileLogger{os.Stdout}

// StderrLogger is a FileLogger that logs to stderr
var StderrLogger = FileLogger{os.Stderr}

// A FileLogger that logs to a given io.Writer
type FileLogger struct {
	W io.Writer
}

func (l *FileLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(l.W, format, v...)
	fmt.Fprintln(l.W)
}

// A BufferLogger simply stores log contents in a bytes.Buffer
type BufferLogger struct {
	Buffer bytes.Buffer
}

func (l *BufferLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&l.Buffer, format, v...)
	fmt.Fprintln(&l.Buffer)
}

func (l *BufferLogger) String() string {
	return l.Buffer.String()
}

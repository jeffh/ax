package ax

import (
	"bytes"
	. "github.com/jeffh/goexpect"
	"testing"
)

func TestFileLogger(t *testing.T) {
	buf := bytes.NewBufferString("")
	l := &FileLogger{buf}
	l.Printf("Hello %s", "world")
	Expect(t, buf.String(), ToEqual, "Hello world\n")
}

func TestBufferLogger(t *testing.T) {
	l := &BufferLogger{}
	l.Printf("Hello %s", "world")
	Expect(t, l.String(), ToEqual, "Hello world\n")
}

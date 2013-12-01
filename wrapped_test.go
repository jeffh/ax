package ax

import (
	. "github.com/jeffh/goexpect"
	"strings"
	"testing"
)

func TestBasicWrappedLogger(t *testing.T) {
	l := &BufferLogger{}
	wrapper := &BasicWrappedLogger{}
	wrapper.SetLogger(l)
	wrapper.Printf("Yo %s", "Dog")
	Expect(t, l.String(), ToEqual, "Yo Dog\n")
	Expect(t, wrapper.WrappedLogger(), ToEqual, l)
}

func TestLockedWrappedLogger(t *testing.T) {
	l := &BufferLogger{}
	wrapper := NewLockedLogger()
	wrapper.SetLogger(l)
	barrier := make(chan bool, 2)
	go func() {
		wrapper.Printf("Hello")
		barrier <- true
	}()
	go func() {
		wrapper.Printf("World")
		barrier <- true
	}()
	<-barrier
	<-barrier
	Expect(t, strings.Contains(l.String(), "Hello"), ToBeTrue)
	Expect(t, strings.Contains(l.String(), "World"), ToBeTrue)
}

func TestPrefixWrappedLogger(t *testing.T) {
	l := &BufferLogger{}
	wrapper := NewPrefixLogger("hi: ")
	wrapper.SetLogger(l)
	wrapper.Printf("lo")
	wrapper.Printf("yo")
	Expect(t, l.String(), ToEqual, "hi: lo\nhi: yo\n")
}

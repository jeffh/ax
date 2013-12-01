Ax
==

A simplier logging interface.

    // The core interface for all logging
    type Logger interface {
        Printf(format string, values ...interface{})
    }

    // The interface for a logger that can wrap another logger
    type WrapLogger interface {
        Logger
        SetLogger(l Logger)
        WrappedLogger() Logger
    }

Usage
=====

Ax provides some shorthand for log management:

    Use(loggers ... Logger) // => Returns a non-nil logger given to Use().
    Use() // => Returns NullLogger
    Use(StdoutLogger) // => Returns stdout logger

In addition, you can wrap loggers to provide more functionality without interupting the logger:

    // returns a Logger that can be written to across goroutines
    Wrap(StdoutLogger, LockedLogger{})

    // returns a Logger that prefixes "[Logger]" to all messages
    Wrap(StdoutLogger, NewPrefixLogger("[Logger]"))

    // combined
    Wrap(StdoutLogger, LockedLogger{}, NewPrefixLogger("[MyPackage]"))

Implementing WrapLogger
-----------------------

For convinence, embed `BasicWrappedLogger`, which implements a No-op interface.
It stores logger in Logger field with Printf passing through.

// Copyright (c) 2018 Claudemiro Alves Feitosa Neto
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package xlog

import (
	"log"
	"os"

	"go.uber.org/fx"
)

type LoggerMapping struct {
	fx.Out

	Logger *log.Logger `name:"x_logger"`
}

// OptionalLogger holds the injected and named 'x_logger' *log.Logger.
// Keep in mind that as it is a optional dependency, it can be nil.
// There is a helper method IsProvided in this same package that you can check if the value was provided or not.
type OptionalLogger struct {
	fx.In

	Logger *log.Logger `name:"x_logger" optional:"true"`
}

// IsProvided returns true if l.OptionalLogger is not nil
func IsProvided(l OptionalLogger) bool {
	return l.Logger != nil
}

// NewLogger returns a new logger configured with values from *Config.
func NewLogger(config Config) LoggerMapping {
	return LoggerMapping{
		Logger: log.New(os.Stdout, config.Prefix, log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
	}
}

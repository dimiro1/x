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

// xtemplate provides html/templates
//
// this module is different from the others because, this one uses a static configuration into the module itself and
// does not load configs from env vars, I have made this decision because usually the location of templates is defined at
// compilation time and not at runtime.
package xtemplate

import (
	"go.uber.org/fx"
)

// Config holds module config
type Config struct {
	RootDir   string
	Extension string
}

// Option used to update config values
type Option func(*Config)

// RootDir option to set the RootDir value on the Config struct.
func RootDir(rootDir string) Option {
	return Option(func(m *Config) {
		m.RootDir = rootDir
	})
}

// Extension option to set the Extension value on the Config struct.
func Extension(ext string) Option {
	return Option(func(m *Config) {
		m.Extension = ext
	})
}

// Module provides a html/template fully configured.
//
// fx.New(xtemplate.Module(xtemplate.RootDir("./templates")))
func Module(options ...Option) fx.Option {
	cfg := &Config{
		RootDir:   "templates",
		Extension: ".html",
	}

	for _, option := range options {
		option(cfg)
	}

	return fx.Options(
		fx.Provide(
			NewTemplate(cfg),
		),
	)
}

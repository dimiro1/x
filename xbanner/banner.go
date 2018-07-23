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

package xbanner

import (
	"context"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"go.uber.org/fx"
)

func Banner(cfg *Config) error {
	in, err := os.Open(cfg.File)
	if in != nil {
		defer in.Close()
	}

	// If the file is not there, just ignore
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return nil
	}

	banner.Init(colorable.NewColorableStdout(), cfg.IsEnabled, cfg.IsColorEnabled, in)
	return nil
}

// registerPrint knows how to show the banner at the right time
func registerPrint(lc fx.Lifecycle, cfg *Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return Banner(cfg)
		},
	})
}

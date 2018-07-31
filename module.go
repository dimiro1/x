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

package x

import (
	"github.com/dimiro1/x/xbanner"
	"github.com/dimiro1/x/xconfig"
	"github.com/dimiro1/x/xhealth"
	"github.com/dimiro1/x/xhttp"
	"github.com/dimiro1/x/xlog"
	"github.com/dimiro1/x/xvars"
	"go.uber.org/fx"
)

// Module x provide an all in one experience, providing a ready to use HTTP Server,
// easy health checks declarations, a few common used middleware.
//
// If you want a more configurable solution, just import the modules that you need to use.
func Module() fx.Option {
	return fx.Options(
		xconfig.Module(),
		xbanner.Module(),
		xlog.Module(),
		xhealth.Module(),
		xhttp.Module(),
		xvars.Module(),
	)
}

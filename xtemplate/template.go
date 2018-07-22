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

package xtemplate

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/fx"
)

// TemplateQualifier is necessary to give a name to the template
type TemplateQualifier struct {
	fx.Out

	Template *template.Template `name:"x_template"`
}

// Template ir necessary to access the template by name
type Template struct {
	fx.In

	Template *template.Template `name:"x_template"`
}

func NewTemplate(config *ModuleConfig) func(funcs FuncMapMappings) (TemplateQualifier, error) {
	return func(funcs FuncMapMappings) (TemplateQualifier, error) {
		funcMap := template.FuncMap{}

		for _, f := range funcs.Functions {
			funcMap[f.Name] = f.Func
		}

		tmpl, err := parseTemplates(config.RootDir, funcMap, config.Extension)

		return TemplateQualifier{
			Template: tmpl,
		}, err
	}
}

// See: https://stackoverflow.com/questions/38686583/golang-parse-all-templates-in-directory-and-subdirectories/38688083
func parseTemplates(rootDir string, funcMap template.FuncMap, ext string) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ext) {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

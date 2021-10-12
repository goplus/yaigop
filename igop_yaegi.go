//go:build yaegi
// +build yaegi

/*
 Copyright 2021 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package igop

import (
	"bytes"
	"errors"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/goplus/gop/cl"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gox"
	"github.com/traefik/yaegi/interp"
)

// -----------------------------------------------------------------------------

func CompileDir(it *interp.Interpreter, srcDir string) (app *interp.Program, err error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, srcDir, nil, 0)
	if err != nil {
		return
	}
	mainPkg, ok := pkgs["main"]
	if !ok {
		return nil, errors.New("not a main package")
	}

	modDir, noCacheFile := findGoModDir(srcDir)
	conf := &cl.Config{
		Dir: modDir, TargetDir: srcDir, Fset: fset, CacheLoadPkgs: true, PersistLoadPkgs: !noCacheFile}
	out, err := cl.NewPackage("", mainPkg, conf)
	if err != nil {
		return
	}
	conf.PkgsLoader.Save()
	var b bytes.Buffer
	err = gox.WriteTo(&b, out, false)
	if err != nil {
		return
	}
	return it.Compile(b.String())
}

/*
// astFileToPkg translate ast.File to ast.Package
func astFileToPkg(file *ast.File, fileName string) (pkg *ast.Package) {
	pkg = &ast.Package{
		Name:  file.Name.Name,
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = file
	return
}
*/

func findGoModFile(dir string) (modfile string, noCacheFile bool, err error) {
	modfile, err = cl.FindGoModFile(dir)
	if err != nil {
		home := os.Getenv("HOME")
		modfile = home + "/gop/go.mod"
		if fi, e := os.Lstat(modfile); e == nil && !fi.IsDir() {
			return modfile, true, nil
		}
		modfile = home + "/goplus/go.mod"
		if fi, e := os.Lstat(modfile); e == nil && !fi.IsDir() {
			return modfile, true, nil
		}
	}
	return
}

func findGoModDir(dir string) (string, bool) {
	modfile, nocachefile, err := findGoModFile(dir)
	if err != nil {
		log.Fatalln("findGoModFile:", err)
	}
	return filepath.Dir(modfile), nocachefile
}

// -----------------------------------------------------------------------------

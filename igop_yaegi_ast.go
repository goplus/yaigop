//go:build yaegiast
// +build yaegiast

package yaigop

import (
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
	file := gox.ASTFile(out, false)
	return it.CompileAST(file)
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

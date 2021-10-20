//go:build ssa
// +build ssa

/*
 Copyright 2020 The GoPlus Authors (goplus.org)
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

package run

import (
	"github.com/goplus/gossa"
	"github.com/goplus/igop"
	"github.com/qiniu/x/log"

	_ "github.com/goplus/gossa/pkg"
)

// -----------------------------------------------------------------------------

func runDir(srcDir string, asm bool) {
	code, err := igop.CompileDir(srcDir)
	if err != nil {
		log.Fatalln(err)
	}
	if asm {
		panic("not impl")
		return
	}
	gossa.RunFile(0, "main.go", code.Data, nil)
}

// -----------------------------------------------------------------------------

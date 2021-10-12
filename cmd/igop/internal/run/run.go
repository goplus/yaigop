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

// Package run implements the ``igop run'' command.
package run

import (
	"fmt"
	"os"
	"runtime"

	"github.com/goplus/gop"
	"github.com/goplus/igop/cmd/igop/internal/base"
	"github.com/qiniu/x/log"
)

// -----------------------------------------------------------------------------

// Cmd - igop run
var Cmd = &base.Command{
	UsageLine: "igop run [-version -quiet -debug] <gopSrcDir>",
	Short:     "Run a Go+ program",
}

var (
	flag        = &Cmd.Flag
	flagAsm     = flag.Bool("asm", false, "generates `asm` code of Go bytecode backend")
	flagQuiet   = flag.Bool("quiet", false, "don't generate any compiling stage log")
	flagDebug   = flag.Bool("debug", false, "print debug information")
	flagVersion = flag.Bool("version", false, "print gop version")
)

func init() {
	Cmd.Run = runCmd
}

func runCmd(cmd *base.Command, args []string) {
	flag.Parse(args)
	if *flagVersion {
		fmt.Println("gop", gop.Version(), runtime.GOOS+"/"+runtime.GOARCH)
		return
	}
	if flag.NArg() < 1 {
		cmd.Usage(os.Stderr)
	}

	log.SetFlags(log.Ldefault &^ log.LstdFlags)
	if *flagQuiet {
		log.SetOutputLevel(0x7000)
	} else if *flagDebug {
		log.SetOutputLevel(log.Ldebug)
	}

	srcDir := flag.Arg(0)
	runDir(srcDir, *flagAsm)
}

// -----------------------------------------------------------------------------

yaigop - Yet Another Go+ interpreter (still in beta version)
========

[![Build Status](https://github.com/goplus/yaigop/actions/workflows/go.yml/badge.svg)](https://github.com/goplus/yaigop/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/goplus/yaigop)](https://goreportcard.com/report/github.com/goplus/yaigop)
[![GitHub release](https://img.shields.io/github/v/tag/goplus/yaigop.svg?label=release)](https://github.com/goplus/yaigop/releases)
[![Playground](https://img.shields.io/badge/playground-Go+-blue.svg)](https://play.goplus.org/)
[![GoDoc](https://pkg.go.dev/badge/github.com/goplus/yaigop.svg)](https://pkg.go.dev/mod/github.com/goplus/yaigop)
[![VSCode](https://img.shields.io/badge/vscode-Go+-teal.svg)](https://github.com/gopcode/vscode-goplus)
[![Interpreter](https://img.shields.io/badge/interpreter-iGo+-seagreen.svg)](https://github.com/goplus/igop)

## Support multiple engines

* [yaegi](https://github.com/traefik/yaegi)
* [ssa](github.com/goplus/gossa)
* [~~gomacro~~](https://github.com/cosmos72/gomacro)
* [~~igo~~](https://github.com/goplus/igo)

## How to build

```bash
git clone git@github.com:goplus/yaigop.git
cd yaigop
go install -tags yaegi -v ./...   # you can replace `yaegi` to `gomacro` or other engines
```

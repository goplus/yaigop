igop - The Go+ interpreter (still in beta version)
========

[![Go Report Card](https://goreportcard.com/badge/github.com/goplus/igop)](https://goreportcard.com/report/github.com/goplus/igop)
[![GitHub release](https://img.shields.io/github/v/tag/goplus/igop.svg?label=release)](https://github.com/goplus/igop/releases)
[![Playground](https://img.shields.io/badge/playground-Go+-blue.svg)](https://play.goplus.org/)
[![VSCode](https://img.shields.io/badge/vscode-Go+-teal.svg)](https://github.com/gopcode/vscode-goplus)
[![GoDoc](https://pkg.go.dev/badge/github.com/goplus/igop.svg)](https://pkg.go.dev/mod/github.com/goplus/igop)

## Support multiple engines

* [yaegi](https://github.com/traefik/yaegi)
* [gomacro](https://github.com/cosmos72/gomacro)
* [ssa](https://github.com/goplus/interp)
* [igo](https://github.com/goplus/igo)

## How to build

```bash
git clone git@github.com:goplus/igop.git
cd igop
go install -tags yaegi -v ./...   # you can replace `yaegi` to `igo` or other engines
```

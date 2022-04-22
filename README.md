igop - The Go+ interpreter (still in beta version)
========

[![Build Status](https://github.com/goplus/igop/actions/workflows/go.yml/badge.svg)](https://github.com/goplus/igop/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/goplus/igop)](https://goreportcard.com/report/github.com/goplus/igop)
[![GitHub release](https://img.shields.io/github/v/tag/goplus/igop.svg?label=release)](https://github.com/goplus/igop/releases)
[![Playground](https://img.shields.io/badge/playground-Go+-blue.svg)](https://play.goplus.org/)
[![VSCode](https://img.shields.io/badge/vscode-Go+-teal.svg)](https://github.com/gopcode/vscode-goplus)
[![GoDoc](https://pkg.go.dev/badge/github.com/goplus/igop.svg)](https://pkg.go.dev/mod/github.com/goplus/igop)

## Support multiple engines

* [ssa](https://github.com/goplus/gossa)
* [yaegi](https://github.com/traefik/yaegi)
* [gomacro](https://github.com/cosmos72/gomacro)
* [igo](https://github.com/goplus/igo)

## How to build

```bash
git clone git@github.com:goplus/igop.git
cd igop
go install -tags ssa -v ./...   # you can replace `ssa` to `yaegi` or other engines
```

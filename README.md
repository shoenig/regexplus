regexplus
=========

Extended `regexp` functions, beyond the standard library

[![Go Report Card](https://goreportcard.com/badge/go.gophers.dev/pkgs/regexplus)](https://goreportcard.com/report/go.gophers.dev/pkgs/regexplus)
[![Build Status](https://travis-ci.com/shoenig/regexplus.svg?branch=master)](https://travis-ci.com/shoenig/regexplus)
[![GoDoc](https://godoc.org/go.gophers.dev/pkgs/regexplus?status.svg)](https://godoc.org/go.gophers.dev/pkgs/regexplus)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/regexplus.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/regexplus.svg)](LICENSE)

# Project Overview

Module `go.gophers.dev/pkgs/regexplus` provides a package with extra functions
for managing regular expressions, beyond what the standard library `regexp` provides.

# Getting Started

The `regexplus` package can be installed by running
```
$ go get go.gophers.dev/pkgs/regexplus
```

#### Example usage
```golang
re := regexp.MustCompile(`(?P<foo>[a-z]+)-(?P<bar>[0-9]+)`)
m := FindNamedSubmatches(re, "abc-123")
// m == Map{"foo":"abc", "bar":"123"}
```

# Contributing

The `go.gophers.dev/pkgs/regexplus` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `go.gophers.dev/pkgs/regexplus` module is open source under the [BSD-3-Clause](LICENSE) license.
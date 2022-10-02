regexplus
=========

Extended `regexp` functions, beyond the standard library

[![GitHub](https://img.shields.io/github/license/shoenig/regexplus.svg)](LICENSE)

# Project Overview

Module `github.com/shoenig/regexplus` provides a package with extra functions
for managing regular expressions, beyond what the standard library `regexp` provides.

# Getting Started

The `regexplus` package can be installed by running
```
$ go get github.com/shoenig/regexplus
```

#### Example usage
```golang
re := regexp.MustCompile(`(?P<foo>[a-z]+)-(?P<bar>[0-9]+)`)
m := FindNamedSubmatches(re, "abc-123")
// m == Map{"foo":"abc", "bar":"123"}
```

# Contributing

The `github.com/shoenig/regexplus` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `github.com/shoenig/regexplus` module is open source under the [BSD-3-Clause](LICENSE) license.
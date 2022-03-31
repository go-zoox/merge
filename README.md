# Merge - Shallow Merge, like JavaScript's Object.assign()

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/merge)](https://pkg.go.dev/github.com/go-zoox/merge)
[![Build Status](https://github.com/go-zoox/merge/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/merge/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/merge)](https://goreportcard.com/report/github.com/go-zoox/merge)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/merge/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/merge?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/merge.svg)](https://github.com/go-zoox/merge/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/merge.svg?label=Release)](https://github.com/go-zoox/merge/releases)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/merge
```

## Getting Started

```go
type Data struct {
  Name string
  Age  int
}

target := Data{Name: "John", Age: 30}
source1 := Data{Name: "Jane", Age: 25}
source2 := Data{Name: "Jack", Age: 35}

Merge(&target, &source1)

// {Name: "Jack", Age: 35}
```

## Depencencies
* [InVisionApp/conjungo](https://github.com/InVisionApp/conjungo) - A small flexible merge library in go.
* [MDN: Object.assign](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/assign) - A JavaScript library for merging objects.

## License
GoZoox is released under the [MIT License](./LICENSE).
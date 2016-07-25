[![Build Status](https://travis-ci.org/markelog/cprf.svg)](https://travis-ci.org/markelog/cprf) [![GoDoc](https://godoc.org/github.com/markelog/cprf?status.svg)](https://godoc.org/github.com/markelog/cprf) [![Go Report Card](https://goreportcard.com/badge/github.com/markelog/cprf)](https://goreportcard.com/report/github.com/markelog/cprf)

# Cprf

`cp -rf <path>` logic on Go

## Installation

```
$ go get github.com/markelog/cprf
```

## Example

```go
package main

import "github.com/markelog/cprf"

func main() {
  // Will extract sexy turtles to current dir
  cprf.Copy("/sexy-turtles", ".")
}
```

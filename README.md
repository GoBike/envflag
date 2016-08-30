# envflag
[![Go Report Card](https://goreportcard.com/badge/github.com/gobike/envflag)](https://goreportcard.com/report/github.com/gobike/envflag)
[![GoDoc](https://godoc.org/github.com/gobike/envflag?status.svg)](https://godoc.org/github.com/gobike/envflag)

Simple environment extension to Golang flag.


## Goals 
- Extends Golang flag with environment-variables.
- Clear precendence: `default` < `environment-variable` < `cli`.

## Installation 

```sh
$ go get github.com/gobike/envflag
```

## Usage

Create main.go

```go
package main

import (
    "fmt"
    "flag"
    "github.com/gobike/envflag"
)

func main() {
    var (
        times int
    )

    flag.IntVar(&times, "f-times", 1, "this is #")
    envflag.Parse() 

    fmt.Println(times)
}
```

Run with `default`.

```sh
$ go run main.go 
1 #output
```

Run with `environment-variable` set.

```sh
$ F_TIMES=100 go run main.go 
100 #output
```

Run with `cli` set.

```sh
$ go run main.go --f-times=10 
10 #output
```

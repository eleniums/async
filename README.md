# async

[![Build Status](https://travis-ci.org/eleniums/async.svg?branch=master)](https://travis-ci.org/eleniums/async) [![Go Report Card](https://goreportcard.com/badge/github.com/eleniums/async)](https://goreportcard.com/report/github.com/eleniums/async) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/eleniums/async/blob/master/LICENSE) [![GoDoc](https://godoc.org/github.com/eleniums/async?status.svg)](https://godoc.org/github.com/eleniums/async)

A collection of methods for running functions concurrently.

## Installation

```
go get -u github.com/eleniums/async
```

## Example

Create some tasks to run:
```go
foo := func() error {
    // do something
    return nil
}

bar := func() error {
    // do something else
    return nil
}
```

Run the tasks concurrently:
```go
err := async.Run(foo, bar)
if err != nil {
    log.Fatalf("task returned an error: %v", err)
}
```

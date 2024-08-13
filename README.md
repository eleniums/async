# async

[![Go Report Card](https://goreportcard.com/badge/github.com/eleniums/async/v2)](https://goreportcard.com/report/github.com/eleniums/async/v2) [![GoDoc](https://godoc.org/github.com/eleniums/async/v2?status.svg)](https://godoc.org/github.com/eleniums/async/v2) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/eleniums/async/blob/master/LICENSE)

A collection of methods for running functions concurrently.

## Installation

```
go get -u github.com/eleniums/async/v2
```

NOTE: `async/v2` is currently identical to `async` as of 6/12/20. When this repo was originally converted to Go Modules, it had already been tagged as `v2.x.x` due to previous breaking changes. At the time, there were some issues converting a repo with major version 2+ to Go Modules without the `/v2` path. As such, the `/v2` path was created to follow convention. Going forward, any changes will be made in the `/v2` path, so only `/v2` should be used.

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
errc := async.Run(foo, bar)
err := async.Wait(errc)
if err != nil {
    log.Fatalf("task returned an error: %v", err)
}
```

# Go Local Cache

[![GoDoc](https://godoc.org/github.com/emahiro/glc?status.svg)](https://godoc.org/github.com/emahiro/glc)
[![Go Report Card](https://goreportcard.com/badge/github.com/emahiro/glc)](https://goreportcard.com/report/github.com/emahiro/glc)

Go Local Cache provides a simple cache mechanism for storing locally.  
Go Local Cache currently only supports in-memory cache, but will also support file cache.

## Installation

```sh
go get github.com/emahiro/glc
```

## Usage

### in memory cache

```go
mc := glc.NewMemoryCache(glc.DefaultMemoryCacheExpires)

// Set
if err := mc.Set("cacheKey", []byte('hoge')); err != nil {
    log.Fatal(err)
}

// Get

data := mc.Get("cacheKey")
```

### file cache

Usage is similar to in memory cache.  
Go Local Cache creates `tmp` directory for file cache.

## LICENSE

MIT

# sl [![Build Status](https://travis-ci.org/frozzare/go-sl.svg?branch=master)](https://travis-ci.org/frozzare/go-sl) [![GoDoc](https://godoc.org/github.com/frozzare/go-sl?status.svg)](https://godoc.org/github.com/frozzare/go-sl) [![Go Report Card](https://goreportcard.com/badge/github.com/frozzare/go-sl)](https://goreportcard.com/report/github.com/frozzare/go-sl)

Go package for dealing with [sl.se](http://sl.se)'s apis.

## Installation

```
$ go get -u github.com/frozzare/go-sl
```

## APIs that are implemented

* [Realtime V4](https://www.trafiklab.se/api/sl-realtidsinformation-4)

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/frozzare/go-sl"
)

func main() {
	api := sl.NewClient(nil)
	res, err := api.Realtime.Search(context.Background(), &sl.RealtimeSearchOptions{
		Key:    "YOUR_API_KEY",
		SiteID: "1002", // 1002 = Stockholm Central/City
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
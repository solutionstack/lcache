# lcache

Cache library for golang. It supports LFU currently.

## Features

* Supports  LFU.

* Goroutine safe.

## Install

```
$ go get github.com/solutionstack/lcache
```

## Example

### Set a key-value pair.

```go
package main

import (
  "github.com/solutionstack/lcache"
  "fmt"
)

func main() {
  lc := lcache.NewCache(20) //optional size parameter to NewCache
  lc.Write("key", "ok")
  
  result := lc.Read("key")
  
  if  result.Error != nil {
    panic( result.Error  )
  }
  fmt.Println("Read:", result.Value)
}

```


# Author

**Olubodn Agbalaya**

* <http://github.com/solutionstack>
* <s.stackng@protonmail.com>

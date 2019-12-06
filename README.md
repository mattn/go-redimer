# go-redimer

Black magic to get slice or array dimmensioned with specified dimensions.

## Usage

```go
package main

import (
    "fmt"

    "github.com/mattn/go-redimer"
)

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    for _, v := range redimer.Redim(&arr, 3, 3).([][3]int) {
        fmt.Println(v)
    }
    // [1 2 3]
    // [4 5 6]
    // [7 8 9]
}

```

## Installation

```
$ go get github.com/mattn/go-redimer
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)

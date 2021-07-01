# go-bencode

## example

```sh
go get github.com/h1zzz/bencode-go
```

```go
package main

import (
    "fmt"

    bencode "github.com/h1zzz/bencode-go"
)

func main() {
    // Decode
    data := []byte("d4:porti6881e1:t2:aa1:y1:e1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:eli201e23:A Generic Error Ocurredee")
    r, err := bencode.Decode(data)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", r)

    // Encode
    v, err := bencode.Encode(r)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", string(v))
}

```

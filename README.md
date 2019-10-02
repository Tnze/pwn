# pwn

Play CTF with golang!

## Getting started
```go
package main

import "github.com/Tnze/pwn"

func main() {
    p := pwn.Remote("example.com:1314")
    p.Write([]byte{0x00, 0x01, 0x02})   // payload
    p.Interactive()
}
```
> There is no `if err != nil`. If an error is present, log.Fatal will be called.
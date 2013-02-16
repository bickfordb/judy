libjudy bindings for go
=======================

libjudy lives here: http://judy.sourceforge.net/

Currently this includes only bindings for the libjudy bitvector implementation since that's all I'm interested in.

Example Usage:

```go

package main

import "judy"

func main() {
  bits := judy.NewJudy1()
  bits.Set(5)
  if bits.Test(5) {
    println("5 exists!")
  }
  for idx, ok := bits.First(0); ok; idx, ok = bits.Next(idx) {
    println(idx)
  }
  // prints 5
}


```



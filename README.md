

```bash
go get github.com/cartersusi/gosimd
```

## Usage 
`main.go`
```go
package main

import (
	"fmt"
	"math/rand"

	simd "github.com/cartersusi/gosimd"
)

func main() {
	n := 200000
	v1 := make([]float32, n)
	v2 := make([]float32, n)

	for i := 0; i < n; i++ {
		v1[i] = rand.Float32()
		v2[i] = rand.Float32()
	}

	var res float32
	simd.DotProduct(v1, v2, res)
	
	fmt.Println(res)
}
```


## Import
```sh
go get github.com/cartersusi/gosimd/avx
```
```go
import "github.com/cartersusi/gosimd/avx"
```

## Usage 
`DotProduct_amd64.go`
```go
import "github.com/cartersusi/gosimd/avx"
func DotProduct(left, right []float32, result float32) (float32, error) {
	if avx.Supported() {
		return avx.DotProduct(left, right, result)
	}
    return errors.New("AVX not supported")
}
```
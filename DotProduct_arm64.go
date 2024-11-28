//go:build arm64

package gosimd

func getSimdImplementation() simdInterface {
	return &fallbackImplementation{}
}

type fallbackImplementation struct{}

func (f *fallbackImplementation) dotProduct(left, right []float32, result *float32) {
	dotProduct(left, right, result)
}

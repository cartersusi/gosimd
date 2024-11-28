//go:build arm64

package gosimd

func getSimdImplementation() SimdInterface {
	return &FallbackImplementation{}
}

type FallbackImplementation struct{}

func (f *FallbackImplementation) _DotProduct(left, right []float32, result *float32) {
	_DotProduct(left, right, result)
}

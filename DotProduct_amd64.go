//go:build amd64

package gosimd

import (
	"github.com/cartersusi/gosimd/avx"
)

func getSimdImplementation() SimdInterface {
	return &AvxImplementation{}
}

type AvxImplementation struct{}

func (a *AvxImplementation) _DotProduct(left, right []float32, result float32) float32 {
	if avx.Supported() {
		return avx.DotProduct(left, right, result)
	}

	return _DotProduct(left, right, result)
}

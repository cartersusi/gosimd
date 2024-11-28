//go:build amd64

package gosimd

import (
	"github.com/cartersusi/gosimd/avx"
)

func getSimdImplementation() SimdInterface {
	return &AvxImplementation{}
}

type AvxImplementation struct{}

func (a *AvxImplementation) _DotProduct(left, right []float32, result *float32) {
	// TODO adjust result param
	if avx.Supported() {
		res := avx.DotProduct(left, right, *result)
		*result = res
		return
	}

	_DotProduct(left, right, result)
}

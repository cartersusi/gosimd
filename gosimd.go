package gosimd

import "runtime"

const (
	SMALL = 65536   //1 << 16
	MED   = 262144  //1 << 18
	LARGE = 1048576 //1 << 20

	OuterBlockSize = 32
	InnerBlockSize = 16
	CacheLineSize  = 64
)

func getOptimalChunkSize(vectorSize int) int {
	numCPU := runtime.GOMAXPROCS(0)
	chunkSize := (vectorSize + numCPU - 1) / numCPU

	const minChunkSize = 1024
	if chunkSize < minChunkSize {
		chunkSize = minChunkSize
	}

	chunkSize = ((chunkSize + 15) / 16) * 16
	if chunkSize > vectorSize {
		chunkSize = vectorSize
	}

	return chunkSize
}

package hash

import (
	"io"
	"sync"
)

// bufferPool is a pool for byte buffers.
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096)
	},
}

// Reader hashes all the bytes in the given reader.
func Reader(reader io.Reader) uint64 {
	var x uint64
	buffer := bufferPool.Get().([]byte)

	for {
		n, err := reader.Read(buffer)
		x = add(x, buffer[:n])

		if err != nil {
			bufferPool.Put(buffer)
			return x
		}
	}
}

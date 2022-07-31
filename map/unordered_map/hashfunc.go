package unordered_map

import (
	"bytes"
	"encoding/gob"
	"hash/maphash"
)

func newHashFunc[K comparable]() func(K) uint64 {
	var h maphash.Hash
	h.SetSeed(maphash.MakeSeed())
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	return func(key K) uint64 {
		_ = encoder.Encode(key)
		_, _ = h.Write(buf.Bytes())
		value := h.Sum64()
		h.Reset()
		buf.Reset()
		return value
	}
}

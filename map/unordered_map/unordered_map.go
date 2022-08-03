package unordered_map

import (
	"bytes"
	"fmt"
	"gontainers/container"
	_map "gontainers/map"
)

type HashMap[K comparable, V any] struct {
	cap int
	len int

	hashFunc func(K) uint64
	table    []*bucket[K, V]
}

// assert HashMap[K,V] to be a container.Container[V]
var _ container.Container[int] = (*HashMap[int, int])(nil)

// assert HashMap[K,V] to be a _map.Map[K,V]
var _ _map.Map[int, int] = (*HashMap[int, int])(nil)

const (
	loadFactorNumerator   = 13
	loadFactorDenominator = 2

	defaultCapacity = 16
)

func New[K comparable, V any]() *HashMap[K, V] {
	capacity := defaultCapacity

	return &HashMap[K, V]{
		cap:      capacity,
		len:      0,
		table:    newTable[K, V](capacity),
		hashFunc: newHashFunc[K](),
	}
}

func newTable[K comparable, V any](size int) []*bucket[K, V] {
	table := make([]*bucket[K, V], size)
	for i := 0; i < size; i++ {
		table[i] = newEmptyBucket[K, V]()
	}
	return table
}

func (h *HashMap[K, V]) Set(key K, value V) {
	hash := h.hashFunc(key)

	b := h.table[hash%uint64(h.cap)]
	if b.set(key, value, hash) {
		h.len++
		if h.len > h.cap*loadFactorNumerator/loadFactorDenominator {
			// TODO: rehash here
			// h.resize()
		}
	}
	panic("not implemented")
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	hash := h.hashFunc(key)

	b := h.table[hash%uint64(h.cap)]
	valuePtr, ok := b.get(key, hash)
	return *valuePtr, ok
}

func (h *HashMap[K, V]) Remove(key K) {
	hash := h.hashFunc(key)

	b := h.table[hash%uint64(h.cap)]
	if b.remove(key, hash) {
		h.len--
	}
}

func (h *HashMap[K, V]) resize() {
	table := newTable[K, V](h.cap * 2)
	_ = table
	for _, b := range h.table {
		_ = b
		// TODO: implement rehash
	}
	panic("not implemented")
}

func (h *HashMap[K, V]) Len() int {
	return h.len
}

func (h *HashMap[K, V]) IsEmpty() bool {
	return h.len == 0
}

func (h *HashMap[K, V]) Keys() []K {
	keys := make([]K, h.len)
	index := 0

	for i := 0; i < h.cap; i++ {
		index = h.table[i].writeKeys(keys, index)
	}
	return keys
}

func (h *HashMap[K, V]) Values() []V {
	values := make([]V, h.len)
	index := 0

	for i := 0; i < h.cap; i++ {
		index = h.table[i].writeValues(values, index)
	}
	return values
}

func (h *HashMap[K, V]) Clear() {
	h.len = 0
	for i := 0; i < h.cap; i++ {
		h.table[i] = newEmptyBucket[K, V]()
	}
	h.hashFunc = newHashFunc[K]()
}

func (h *HashMap[K, V]) String() string {
	var mask uint64
	buf := bytes.NewBufferString("{")

	for _, b := range h.table {
		for b != nil {
			mask = (1 << 8) - 1
			for i := 0; i < 8; i++ {
				if b.reserved&mask == 1 {
					buf.WriteString(fmt.Sprintf("%v:%v, ", b.keys[i], b.values[i]))
				}
				mask <<= 8
			}
			b = b.next
		}
	}
	buf.Truncate(buf.Len() - 2)
	buf.WriteString("}")
	return buf.String()
}

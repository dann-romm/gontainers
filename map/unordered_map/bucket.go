package unordered_map

type bucket[K comparable, V any] struct {
	next     *bucket[K, V]
	reserved uint64 // 0: empty, 1: occupied
	bitmask  uint64
	keys     [8]K
	values   [8]V
}

func newEmptyBucket[K comparable, V any]() *bucket[K, V] {
	return &bucket[K, V]{
		next:     nil,
		reserved: 0,
		bitmask:  0,
		keys:     [8]K{},
		values:   [8]V{},
	}
}

func newBucket[K comparable, V any](key K, value V, hash uint64) *bucket[K, V] {
	return &bucket[K, V]{
		next:     nil,
		reserved: 1,
		bitmask:  hash & ((1 << 8) - 1),
		keys:     [8]K{key},
		values:   [8]V{value},
	}
}

func (b *bucket[K, V]) get(key K, hash uint64) (*V, bool) {
	var mask uint64 = (1 << 8) - 1

	for i := 0; i < 8; i++ {
		if b.reserved&mask == 1 && b.bitmask&mask == hash&mask && b.keys[i] == key {
			return &b.values[i], true
		}
		mask <<= 8
	}

	if b.next == nil {
		return nil, false
	}
	return b.next.get(key, hash)
}

func (b *bucket[K, V]) set(key K, value V, hash uint64) bool {
	if valuePtr, ok := b.get(key, hash); ok {
		*valuePtr = value
		return false
	}
	b.setNew(key, value, hash)
	return true
}

func (b *bucket[K, V]) setNew(key K, value V, hash uint64) {
	var mask uint64 = (1 << 8) - 1

	for i := 0; i < 8; i++ {
		if b.reserved&mask == 0 {
			b.reserved |= 1 << i * 8
			b.bitmask |= hash & mask
			b.keys[i] = key
			b.values[i] = value
			return
		}
		mask <<= 8
	}
	if b.next == nil {
		b.next = newBucket(key, value, hash)
		return
	}
	b.next.setNew(key, value, hash)
}

func (b *bucket[K, V]) remove(key K, hash uint64) bool {
	var mask uint64 = (1 << 8) - 1

	for i := 0; i < 8; i++ {
		if b.reserved&mask == 1 && b.bitmask&mask == hash&mask && b.keys[i] == key {
			b.reserved &= ^mask
			b.bitmask &= ^mask
			return true
		}
		mask <<= 8
	}
	if b.next == nil {
		return false
	}
	return b.next.remove(key, hash)
}

func (b *bucket[K, V]) writeKeys(keys []K, index int) int {
	var mask uint64 = (1 << 8) - 1

	for i := 0; i < 8; i++ {
		if b.reserved&mask == 1 {
			keys[index] = b.keys[i]
			index++
		}
		mask <<= 8
	}
	if b.next == nil {
		return index
	}
	return b.next.writeKeys(keys, index)
}

func (b *bucket[K, V]) writeValues(values []V, index int) int {
	var mask uint64 = (1 << 8) - 1

	for i := 0; i < 8; i++ {
		if b.reserved&mask == 1 {
			values[index] = b.values[i]
			index++
		}
		mask <<= 8
	}
	if b.next == nil {
		return index
	}
	return b.next.writeValues(values, index)
}

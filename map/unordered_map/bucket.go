package unordered_map

type bucket[K comparable, V any] struct {
	reserved uint64 // 0: empty, 1: occupied
	bitmask  uint64
	keys     [8]K
	values   [8]V
	next     *bucket[K, V]
}

func newEmptyBucket[K comparable, V any]() *bucket[K, V] {
	return &bucket[K, V]{
		reserved: 0,
		bitmask:  0,
		keys:     [8]K{},
		values:   [8]V{},
		next:     nil,
	}
}

func newBucket[K comparable, V any](key K, value V, hash uint64) *bucket[K, V] {
	return &bucket[K, V]{
		reserved: 1,
		bitmask:  hash & ((1 << 8) - 1),
		keys:     [8]K{key},
		values:   [8]V{value},
		next:     nil,
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

func (b *bucket[K, V]) set(key K, value V, hash uint64) {
	if valuePtr, ok := b.get(key, hash); ok {
		*valuePtr = value
		return
	}
	b.setNew(key, value, hash)
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

func (b *bucket[K, V]) remove(key K, hash uint64) {
	var mask uint64 = (1 << 8) - 1

	for i := 0; i < 8; i++ {
		if b.reserved&mask == 1 && b.bitmask&mask == hash&mask && b.keys[i] == key {
			b.reserved &= ^mask
			b.bitmask &= ^mask
			return
		}
		mask <<= 8
	}
	if b.next == nil {
		return
	}
	b.next.remove(key, hash)
}

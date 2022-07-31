package unordered_map

type HashMap[K comparable, V any] struct {
	cap        int
	len        int
	reserved   int
	loadFactor float64

	hashFunc func(K) uint64
	table    []*bucket[K, V]
}

const (
	defaultLoadFactor = 0.75
	defaultCapacity   = 16
)

func New[K comparable, V any]() *HashMap[K, V] {
	capacity := defaultCapacity

	return &HashMap[K, V]{
		cap:        capacity,
		len:        0,
		reserved:   0,
		loadFactor: defaultLoadFactor,
		table:      newTable[K, V](capacity),
		hashFunc:   newHashFunc[K](),
	}
}

func newTable[K comparable, V any](size int) []*bucket[K, V] {
	table := make([]*bucket[K, V], size)
	for i := 0; i < size; i++ {
		table[i] = newEmptyBucket[K, V]()
	}
	return table
}

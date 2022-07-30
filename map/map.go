package _map

import "gontainers/container"

type Map[K comparable, V any] interface {
	container.Container[V]

	// Set adds or updates the value for the given key
	Set(key K, value V)

	// Get returns the value for the given key
	Get(key K) (V, bool)

	// Remove removes the value for the given key
	Remove(key K)

	// Keys returns slice of all keys in the map
	Keys() []K
}

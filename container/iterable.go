package container

type ForwardIterable[T any] interface {
	// Begin returns an iterator to the first element in the container
	Begin() Iterator[T]
}

type ReverseIterable[T any] interface {
	// RBegin returns a reverse iterator to the first element in the container
	RBegin() Iterator[T]
}

type Iterable[T any] interface {
	ForwardIterable[T]
	ReverseIterable[T]
}

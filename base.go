package gontainer

// New returns a in-memory container, intended for prototyping and testing.
func New[K comparable, V any]() Container[K, V] {
	return make(mapWrap[K, V])
}

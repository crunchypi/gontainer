# gontainer
Generic container interfaces and decorators for Go

Index
- [Core interfaces](#core-interfaces)
- [Errors](#errors)
- [Impl pattern](#impl-pattern)
- [Default](#default)



## Core interfaces
Core interfaces are just CRUD operations & search-based operations.

#### Basic CRUD.
```go
type Putter[K comparable, V any] interface {
	Put(ctx context.Context, key K, val V) (err error)
}

type Getter[K comparable, V any] interface {
	Get(ctx context.Context, key K) (val V, err error)
}

type Modifier[K comparable, V any] interface {
	Mod(ctx context.Context, key K, rcv func(v V) V) (err error)
}

type Deleter[K comparable, V any] interface {
	Del(ctx context.Context, key K) (val V, err error)
}
```

#### Searcher-based.
```go
type Searcher[Q, R any] interface {
	Search(ctx context.Context, filter Q) (r R, err error)
}

type SearchUpdater[Q, U, R any] interface {
	SearchUpdate(ctx context.Context, filter Q, update U) (r R, err error)
}

type SearchDeleter[Q, R any] interface {
	SearchDelete(ctx context.Context, filter Q) (r R, err error)
}
```

#### Composite/Container
```go
type Container[K comparable, V any] interface {
	Putter[K, V]
	Getter[K, V]
	Modifier[K, V]
	Deleter[K, V]

	Len(context.Context) (int, error)
	Cap(context.Context) (int, error)
}
```



## Errors
```go
var ErrPut = errors.New("gontainer: failed put")
var ErrGet = errors.New("gontainer: failed get")
var ErrMod = errors.New("gontainer: failed mod")
var ErrDel = errors.New("gontainer: failed del")

var ErrSearchFinder = errors.New("gontainer: failed search")
var ErrSearchUpdater = errors.New("gontainer: failed search & update")
var ErrSearchDeleter = errors.New("gontainer: failed search & update")

// See the next section.
var ErrImpl = errors.New("gontainer: used interface without an implementation")
```



## Impl pattern

The impl pattern allows you to implement an interface in a functional way, avoiding the tedium of defining structs which implement small interfaces. You simply define a function with the right signature and place it inside an impl struct.




## Default
A default container is implemented and has the signature noted below. It is mainly intended for prototyping and testing, and is implemented as a `map[K]V`wrapped by the `Container` interface. The underlying implementation may be swapped in the future but the signature and behavior will most likely not.

Some notes:
- As `cap(map[K]V)` is not supported by the language, a call to `Cap` returns `Len` * 2
- `Mod`will run the callback and save the result even if the key does not exist.

```go
func New[K comparable, V any]() Container[K, V]
```
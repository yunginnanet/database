# database

[![Coverage](https://codecov.io/gh/yunginnanet/database/branch/master/graph/badge.svg)](https://codecov.io/gh/yunginnanet/database)
[![Build Status](https://github.com/yunginnanet/database/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/yunginnanet/database/actions/workflows/go.yml)

---

#### type Filer

```go
type Filer interface {

	// Has should return true if the given key has an associated value.
	Has(key []byte) bool
	// Get should retrieve the byte slice corresponding to the given key, and any associated errors upon failure.
	Get(key []byte) ([]byte, error)
	// Put should insert the value data in a way that is associated and can be retrieved by the given key data.
	Put(key []byte, value []byte) error
	// Delete should delete the key and the value associated with the given key, and return an error upon failure.
	Delete(key []byte) error
}
```

Filer is is a way to implement any generic key/value store. These functions
should be plug and play with most of the popular key/value store golang
libraries.

#### type Keeper

```go
type Keeper interface {
	// Path should return the base path where all buckets should be stored under. (likely as subdirectories)
	Path() string
	// Init should initialize our Filer at the given path, to be referenced and called by bucketName.
	Init(bucketName string) error
	// With provides access to the given bucketName by providing a pointer to the related Filer.
	With(bucketName string) Filer
	// Close should safely end any Filer operations of the given bucketName and close any relevant handlers.
	Close(bucketName string) error
	// Sync should take any volatile data and solidify it somehow if relevant. (ram to disk in most cases)
	Sync(bucketName string) error

	CloseAll() error
	SyncAll() error
}
```

Keeper will be in charge of the more meta operations involving Filers. This
includes operations like initialization, syncing to disk if applicable, and
backing up.

NOTE: Many key/value golang libraries will already implement this interface
already. This exists for more potential granular control in the case that they
don't. Otherwise you'd have to build a wrapper around an existing key/value
store to satisfy an overencompassing interface.

#### type Key

```go
type Key interface {
	Bytes() []byte
	String() string
	Equal(Key) bool
}
```

Key represents a key in a key/value Filer.

#### type Searcher

```go
type Searcher interface {
	// AllKeys must retrieve all keys in the datastore with the given bucketName.
	AllKeys() []string
	// PrefixScan must return all keys that begin with the given prefix.
	PrefixScan(prefix string) map[string]interface{}
	// Search must be able to search through the contents of our database and return a map of results.
	Search(query string) map[string]interface{}
	// ValueExists searches for an exact match of the given value and returns the key that contains it.
	ValueExists(value []byte) (key []byte, ok bool)
}
```

Searcher must be able to search through our datastore(s) with strings.

#### type Value

```go
type Value interface {
	Bytes() []byte
	String() string
	Equal(Value) bool
}
```

Value represents a value in a key/value Filer.

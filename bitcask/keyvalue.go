package bitcask

import (
	"git.tcp.direct/kayos/common/hash"

	"git.tcp.direct/tcp.direct/database"
)

// KeyValue represents a key and a value from a key/balue store.
type KeyValue struct {
	Key   Key
	Value Value
}

// Key represents a key in a key/value store.
type Key struct {
	database.Key
	b []byte
}

// Bytes returns the raw byte slice form of the Key.
func (k Key) Bytes() []byte {
	return k.b
}

// String returns the string slice form of the Key.
func (k Key) String() string {
	return string(k.b)
}

// Equal determines if two keys are equal.
func (k Key) Equal(k2 Key) bool {
	return hash.BlakeEqual(k.Bytes(), k2.Bytes())
}

// Value represents a value in a key/value store.
type Value struct {
	database.Value
	b []byte
}

// Bytes returns the raw byte slice form of the Value.
func (v Value) Bytes() []byte {
	return v.b
}

// String returns the string slice form of the Value.
func (v Value) String() string {
	return string(v.b)
}

// Equal determines if two values are equal.
func (v Value) Equal(v2 Value) bool {
	return hash.BlakeEqual(v.Bytes(), v2.Bytes())
}

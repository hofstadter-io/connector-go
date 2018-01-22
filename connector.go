package connector

import "reflect"

// Named ...
//   Things which have a Name()
type Named interface {
	Name() string
}

// Itemizer ...
//   Things which have Items()
type Itemizer interface {
	Items() []interface{}
}

// Connected ...
//   Things which have Connectors()
type Connected interface {
	Connectors() []Connector
}

// Addable ...
//   Things which can be Add()-ed to
type Addable interface {
	Add(...interface{})
}

// Deletable ...
//   Things which can be Del()-ed from
type Deletable interface {
	Del(interface{})
}

// Clearable ...
//   Things which are Clear()-able
type Clearable interface {
	Clear()
}

// Stored ...
//   Things that hold other things [Add(),Del(),Clear()]
type Stored interface {
	Addable
	Deletable
	Clearable
}

// Indexed ...
//   Things that have something to Lookup(string)
type Indexed interface {
	Lookup(name string) interface{}
}

// Searchable ...
//   Things that have something to Find(interface)
type Searchable interface {
	Find(interface{})
}

// Filterable ...
//   Things that are Filter(func(interface{})bool)-able
type Filterable interface {
	Filter(func(interface{}) bool) []interface{}
}

// Connector ...
//   is all the things, put together. Go Get(reflect.Type)'em
type Connector interface {
	Named
	Itemizer
	Connected
	Addable
	// Stored
	// Indexed
	// Searchable

	Get(reflect.Type) []interface{}
}

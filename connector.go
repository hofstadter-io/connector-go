package connector

import "reflect"

// Named - things which have a Name()
type Named interface {
	Name() string
}

// Namer - things which hold things which have a Name()
type Namer interface {
	Named() []Named
}

// Itemizer - things which have Items()
type Itemizer interface {
	Items() []interface{}
}

// Connected - things which have Connectors()
type Connected interface {
	Connectors() []Connector
}

// Addable - things which can be Add()-ed to
type Addable interface {
	Add(...interface{})
}

// Deletable - things which can be Del()-ed from
type Deletable interface {
	Del(interface{})
}

// Clearable - things which are Clear()-able
type Clearable interface {
	Clear()
}

// Stored - hings that hold other things [Add(),Del(),Clear()]
type Stored interface {
	Addable
	Deletable
	Clearable
}

// Indexed - things that have something to Lookup(string)
type Indexed interface {
	Lookup(name string) interface{}
}

// Searchable - things that have something to Find(interface)
type Searchable interface {
	Find(interface{})
}

// Filterable - things that are Filter(func(interface{})bool)-able
type Filterable interface {
	Filter(func(interface{}) bool) []interface{}
}

// Connector is all the things, put together. Go Get(reflect.Type)'em
type Connector interface {
	Named

	Get(reflect.Type) []interface{}

	Namer
	Itemizer
	Connected
	Addable
	// Stored
	// Indexed
	// Searchable

}

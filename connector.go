package connector

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

// Gettable - things which can be Get()-ed from.
// Extracts all the items which match a given type. Go Get()'em !!
type Gettable interface {
	Get(interface{}) []interface{}
}

// Deletable - things which can be Del()-ed from
type Deletable interface {
	Del(interface{})
}

// Clearable - things which are Clear()-able
type Clearable interface {
	Clear()
}

// Stored - hings that hold other things [Add(),Get(),Del(),Clear()]
type Stored interface {
	Addable
	Gettable
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

// Connector is all the things, put together.
//
// Connect() recursively passes a Connector object to all items
// so that they may consume or use any items in the Connector.
// Typically, we build up a root Connector and then
// call Connect with itself as the argument.
type Connector interface {
	Connect(Connector)

	Named
	Namer
	Itemizer
	Connected
	Gettable
	Addable
	// Deletable
	// Clearable
	// Indexed
	// Searchable
	// Filterable

}

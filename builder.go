package connector

import (
	"fmt"
	"reflect"
)

// Builder ...
//   is the 'Connector' implementation
type Builder struct {
	name  string
	items []interface{}
}

// New ...
// Creates a new Builder (i.e. Connector) with the given name.
// Name must have a non-emtpy value.
// It should also be unique, but that is not enforced.
func New(name string) *Builder {
	if name == "" {
		return nil
	}
	return &Builder{
		name:  name,
		items: []interface{}{},
	}
}

// Name ...
// Returns the name of the Builder (i.e. Connector)
func (B *Builder) Name() string {
	return B.name
}

// Items ...
// Returns all Items
func (B *Builder) Items() []interface{} {
	all := []interface{}{}
	for _, item := range B.items {
		all = append(all, item)

		// recurse if (also) module
		itmzr, ok := item.(Itemizer)
		if ok {
			all = append(all, itmzr.Items()...)
		}
	}

	return all
}

// Connectors ...
// Returns all Connector's found amongst the items
func (B *Builder) Connectors() []Connector {
	all := []Connector{}
	for _, item := range B.items {
		conn, ok := item.(Connector)
		if ok {
			all = append(all, conn)
			all = append(all, conn.Connectors()...)
		}
	}

	return all
}

// Add ...
// Adds items to a Connector. May be a single object, slice, or any mix.
func (B *Builder) Add(in ...interface{}) {
	B.add(in)
}

func (B *Builder) add(in interface{}) {
	switch it := in.(type) {
	case []interface{}:
		for _, i := range it {
			B.add(i)
		}
	default:
		B.items = append(B.items, it)

		itmzr, ok := in.(Itemizer)
		if ok {
			for _, i := range itmzr.Items() {
				B.add(i)
			}
		}
	}
}

// Del ...
// Deletes an item or list of items from the connector.
func (B *Builder) Del(out interface{}) {
	switch ot := out.(type) {
	case []interface{}:
		for _, o := range ot {
			B.Del(o)
		}
		return
	}

	// TODO otherwise look for objects
}

// Clear ...
// Clears all items from this Connector.
// Will not effect connectors which this had been added to.
func (B *Builder) Clear() {
	B.items = []interface{}{}
}

/*Get ...

Get is the main function of Connector.

Once a number of items have been added,
we will want to retrieve some subset of those.
The Get() Function will return all items
that match the supplied type

*/
func (B *Builder) Get(typ reflect.Type) []interface{} {
	all := []interface{}{}

	for _, item := range B.items {
		it := reflect.TypeOf(item)
		// iel := it.Elem()
		// tel := typ.Elem()
		fmt.Println(typ, it)
		// fmt.Println(iel, tel)
		fmt.Println()

		continue
		/*
			if typ.Implements(iel) {
				all = append(all, item)
			}
			conn, ok := item.(Connector)
			if ok {
				all = append(all, conn.Get(typ)...)
			}
		*/
	}

	return all
}

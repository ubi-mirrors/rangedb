// Code generated by go generate
package rangedbtest

import "github.com/inklabs/rangedb"

type eventBinder interface {
	Bind(events ...rangedb.Event)
}

func BindEvents(binder eventBinder) {
	binder.Bind(
		&ThingWasDone{},
		&AnotherWasComplete{},
		&ThatWasDone{},
	)
}
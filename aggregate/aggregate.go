package aggregate

type AggregateRoot interface {
	Entity
	Version() int
}

type Entity interface {
	Events() EventStream
	AddEvents(events ...interface{})
	ClearEvents()
}

type aggregateRoot struct {
	version int
	events  EventStream
}

func NewAggregateRoot(version int) AggregateRoot {
	return &aggregateRoot{version: version}
}

func (a *aggregateRoot) Events() EventStream {
	return a.events
}

func (a *aggregateRoot) AddEvents(events ...interface{}) {
	a.events = append(a.events, events...)
}

func (a *aggregateRoot) ClearEvents() {
	if len(a.events) >= 1 {
		a.events = EventStream{}
		a.version++
	}
}

func (a *aggregateRoot) Version() int {
	return a.version
}

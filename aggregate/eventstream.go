package aggregate

import (
	"github.com/mbict/go-eventbus"
	"reflect"
)

type EventStream []interface{}

func (e *EventStream) IsEmpty() bool {
	return e == nil || len(*e) == 0
}

func (e EventStream) HasAnyOf(event interface{}) bool {
	if hasEvt, ok := event.(eventbus.Event); ok {
		for _, v := range e {
			if evt, ok := v.(eventbus.Event); ok && evt.EventType() == hasEvt.EventType() {
				return true
			}
		}
	} else { //fallback on reflection
		et := reflect.TypeOf(event)
		for _, v := range e {
			if reflect.TypeOf(v) == et {
				return true
			}
		}
	}
	return false
}

func (e EventStream) GetOf(event interface{}) []interface{} {
	var res []interface{}

	if hasEvt, ok := event.(eventbus.Event); ok {
		for _, v := range e {
			if evt, ok := v.(eventbus.Event); ok && evt.EventType() == hasEvt.EventType() {
				res = append(res, v)
			}
		}
	} else { //fallback on reflection
		et := reflect.TypeOf(event)
		for _, v := range e {
			if reflect.TypeOf(v) == et {
				res = append(res, v)
			}
		}
	}

	return res
}

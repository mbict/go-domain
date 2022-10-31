package aggregate

import (
	"github.com/mbict/go-domain/v2"
	"reflect"
)

type EventStream []any

func (e *EventStream) IsEmpty() bool {
	return e == nil || len(*e) == 0
}

func (e EventStream) HasAnyOf(event any) bool {
	if hasEvt, ok := event.(domain.Event); ok {
		for _, v := range e {
			if evt, ok := v.(domain.Event); ok && evt.EventName() == hasEvt.EventName() {
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

func (e EventStream) GetOf(event any) []any {
	var res []any

	if hasEvt, ok := event.(domain.Event); ok {
		for _, v := range e {
			if evt, ok := v.(domain.Event); ok && evt.EventName() == hasEvt.EventName() {
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

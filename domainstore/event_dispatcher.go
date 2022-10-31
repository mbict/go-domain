package domainstore

import (
	"context"
	"github.com/google/uuid"
	"github.com/mbict/go-domain/v2/aggregate"
)

type AggregateStore[T any] interface {
	Load(ctx context.Context, id uuid.UUID) (T, error)
	Store(ctx context.Context, aggregate T) error
}

type DispatcherFunc func(stream aggregate.EventStream)

type eventDispatchingAggregateStore[T any] struct {
	aggregateStore AggregateStore[T]
	dispatchFunc   DispatcherFunc
}

func (s *eventDispatchingAggregateStore[T]) Load(ctx context.Context, id uuid.UUID) (T, error) {
	return s.aggregateStore.Load(ctx, id)
}

func (s *eventDispatchingAggregateStore[T]) Store(ctx context.Context, aggr T) error {
	aggregateRoot := any(aggr).(aggregate.AggregateRoot)
	eventStream := aggregateRoot.Events()

	if err := s.aggregateStore.Store(ctx, aggr); err != nil {
		return err
	}

	s.dispatchFunc(eventStream)
	return nil
}

func NewEventDispatchingAggregateStore[T any](aggregateStore AggregateStore[T], dispatchFunc DispatcherFunc) AggregateStore[T] {
	return &eventDispatchingAggregateStore[T]{
		aggregateStore: aggregateStore,
		dispatchFunc:   dispatchFunc,
	}
}

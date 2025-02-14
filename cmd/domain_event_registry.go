package main

import (
	"reflect"

	"github.com/renatocosta55sp/event_sourcing/internal/events"
	"github.com/renatocosta55sp/modeling/infra/bus"
)

func configureDomainEventRegistry() *bus.EventRegistry {

	eventRegistry := bus.NewEventRegistry()
	eventRegistry.RegisterEvents(map[string]reflect.Type{
		(&events.AccountOpened{}).GetName():      reflect.TypeOf(events.AccountOpened{}),
		(&events.FundsDeposited{}).GetName():     reflect.TypeOf(events.FundsDeposited{}),
		(&events.FundsWithdrawn{}).GetName():     reflect.TypeOf(events.FundsWithdrawn{}),
		(&events.AccountMonthOpened{}).GetName(): reflect.TypeOf(events.AccountMonthOpened{}),
		(&events.AccountMonthClosed{}).GetName(): reflect.TypeOf(events.AccountMonthClosed{}),
	})

	return eventRegistry

}

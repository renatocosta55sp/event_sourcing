package removedevice

import (
	"context"

	"github.com/renatocosta55sp/device_management/internal/domain"
	"github.com/renatocosta55sp/device_management/internal/domain/commands"
	"github.com/renatocosta55sp/device_management/internal/infra/adapters/persistence"
	"github.com/renatocosta55sp/modeling/eventstore"
	"github.com/renatocosta55sp/modeling/infra/bus"
	"github.com/renatocosta55sp/modeling/slice"
)

type CommandExecutor struct {
	eventStore  eventstore.EventStore
	snapshot    eventstore.SnapshotStore
	transaction persistence.TransactionDb
}

func (c CommandExecutor) Send(ctx context.Context, cmd commands.RemoveDeviceCommand) (commandResult slice.CommandResult, device *domain.DeviceAggregate, err error) {

	err = c.transaction.Transaction(func() (err error) {

		//Get the current state
		version, err := c.snapshot.ReadSnapshot(ctx, cmd.AggregateID.String())
		if err != nil {
			return err
		}

		stream, err := c.eventStore.ReadStream(ctx, cmd.AggregateID.String(), version)
		if err != nil {
			return err
		}

		deviceAggregate := domain.NewDeviceAggregate(stream)

		commandResult, err = deviceAggregate.Remove(cmd)
		if err != nil {
			return err
		}

		dispatcher := bus.NewEventDispatcher()

		deviceReadModel := DeviceReadModel{deviceAggregate: deviceAggregate, eventStore: c.eventStore, snapshot: c.snapshot, ctx: ctx}
		bus.RegisterHandler(dispatcher, deviceReadModel)

		if err := dispatcher.DispatchUncommittedEvents(deviceAggregate.UncommittedEvents); err != nil {
			return err
		}

		device = deviceAggregate

		return nil
	})

	return commandResult, device, err
}

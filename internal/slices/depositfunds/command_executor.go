package depositfunds

import (
	"context"

	"github.com/renatocosta55sp/event_sourcing/internal/domain"
	"github.com/renatocosta55sp/event_sourcing/internal/domain/commands"
	"github.com/renatocosta55sp/event_sourcing/internal/infra/adapters/persistence"
	"github.com/renatocosta55sp/modeling/eventstore"
	"github.com/renatocosta55sp/modeling/infra/bus"
	"github.com/renatocosta55sp/modeling/slice"
)

type CommandExecutor struct {
	eventStore  eventstore.EventStore
	snapshot    eventstore.SnapshotStore
	transaction persistence.TransactionDb
}

func (c CommandExecutor) Send(ctx context.Context, cmd commands.DepositFundsCommand) (commandResult slice.CommandResult, bankAccount *domain.BankAccountAggregate, err error) {

	err = c.transaction.Transaction(func() (err error) {

		version, err := c.snapshot.ReadSnapshot(ctx, cmd.AggregateID.String())
		if err != nil {
			return err
		}

		stream, err := c.eventStore.ReadStream(ctx, cmd.AggregateID.String(), version)

		if err != nil {
			return err
		}

		bankAccountAggregate := domain.NewBankAccountAggregate(stream)

		commandResult, err = bankAccountAggregate.Deposit(cmd)
		if err != nil {
			return err
		}

		dispatcher := bus.NewEventDispatcher()

		deviceReadModel := BankAccountReadModel{bankAccountAggregate: bankAccountAggregate, eventStore: c.eventStore, snapshot: c.snapshot, ctx: ctx}
		bus.RegisterHandler(dispatcher, deviceReadModel)

		if err := dispatcher.DispatchUncommittedEvents(bankAccountAggregate.UncommittedEvents); err != nil {
			return err
		}

		bankAccount = bankAccountAggregate

		return nil

	})

	return commandResult, bankAccount, err
}

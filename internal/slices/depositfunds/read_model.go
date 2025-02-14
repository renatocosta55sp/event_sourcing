package depositfunds

import (
	"context"

	"github.com/renatocosta55sp/event_sourcing/internal/domain"
	"github.com/renatocosta55sp/event_sourcing/internal/events"
	"github.com/renatocosta55sp/modeling/eventstore"
)

type BankAccountReadModel struct {
	bankAccountAggregate *domain.BankAccountAggregate
	eventStore           eventstore.EventStore
	snapshot             eventstore.SnapshotStore
	ctx                  context.Context
}

func (b BankAccountReadModel) Handle(event *events.FundsDeposited) error {

	if err := b.eventStore.AppendToStream(b.ctx,
		event.AggregateId.String(),
		b.bankAccountAggregate.UncommittedEvents,
		b.bankAccountAggregate.Version); err != nil {
		return err
	}

	if eventstore.ShouldTakeSnapshot(&b.bankAccountAggregate.Aggregate, 5) {
		if err := b.snapshot.WriteSnapshot(b.ctx, event.AggregateId.String(), event, b.bankAccountAggregate.Version); err != nil {
			return err
		}
	}

	return nil
}

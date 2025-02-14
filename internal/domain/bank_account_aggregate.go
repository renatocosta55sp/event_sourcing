package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/renatocosta55sp/event_sourcing/internal/domain/commands"
	"github.com/renatocosta55sp/event_sourcing/internal/events"
	"github.com/renatocosta55sp/modeling/domain"
	"github.com/renatocosta55sp/modeling/slice"
)

type BankAccountAggregate struct {
	domain.Aggregate
	Balance float64
}

func NewBankAccountAggregate(stream []domain.Event) *BankAccountAggregate {
	d := &BankAccountAggregate{}
	d.hydrate(stream)
	return d
}

func (d *BankAccountAggregate) hydrate(stream []domain.Event) {

	lenStream := len(stream)

	if lenStream > 0 {
		for _, e := range stream {
			d.Apply(e)
		}
		d.Version = stream[lenStream-1].GetVersion()
	}

}

func (d *BankAccountAggregate) Apply(event domain.Event) {

	switch e := event.(type) {

	case *events.FundsDeposited:
		d.Balance = e.NewBalance

	case *events.FundsWithdrawn:
		d.Balance = e.NewBalance
	}

}

func (d *BankAccountAggregate) Deposit(cmd commands.DepositFundsCommand) (slice.CommandResult, error) {

	commandResult := slice.CommandResult{
		Identifier:        cmd.AggregateID,
		AggregateSequence: d.Version,
	}

	if cmd.Amount <= 0 {
		return commandResult, ErrAmountMustBePositive
	}

	event := &events.FundsDeposited{
		AggregateId: cmd.AggregateID,
		Amount:      cmd.Amount,
		NewBalance:  d.Balance + cmd.Amount,
	}

	event.SetId(uuid.New())

	event.SetVersion(d.Version + 1)

	d.UncommittedEvents = append(d.UncommittedEvents, event)
	d.Apply(event)

	d.Version++

	return commandResult, nil

}

func (d *BankAccountAggregate) Withdraw(cmd commands.WithdrawnFundsCommand) (slice.CommandResult, error) {

	commandResult := slice.CommandResult{
		Identifier:        cmd.AggregateID,
		AggregateSequence: d.Version,
	}

	if cmd.Amount <= 0 {
		return commandResult, ErrAmountMustBePositive
	}

	event := &events.FundsWithdrawn{
		AggregateId: cmd.AggregateID,
		Amount:      cmd.Amount,
		NewBalance:  d.Balance - cmd.Amount,
	}

	event.SetId(uuid.New())
	event.SetVersion(d.Version + 1)

	d.UncommittedEvents = append(d.UncommittedEvents, event)

	d.Apply(event)

	d.Version++

	return commandResult, nil

}

var (
	ErrAmountMustBePositive = errors.New("Deposit amount must be positive")
)

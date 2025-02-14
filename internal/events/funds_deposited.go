package events

import "github.com/google/uuid"

type FundsDeposited struct {
	Id, AggregateId    uuid.UUID
	Amount, NewBalance float64
	Version            int
}

func (e *FundsDeposited) GetName() string {
	return "FundsDeposited"
}

func (e *FundsDeposited) GetId() uuid.UUID {
	return e.Id
}

func (e *FundsDeposited) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *FundsDeposited) GetVersion() int {
	return e.Version
}

func (e *FundsDeposited) SetVersion(number int) {
	e.Version = number
}

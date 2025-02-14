package events

import "github.com/google/uuid"

type FundsWithdrawn struct {
	Id, AggregateId    uuid.UUID
	Amount, NewBalance float64
	Version            int
}

func (e *FundsWithdrawn) GetName() string {
	return "FundsWithdrawn"
}

func (e *FundsWithdrawn) GetId() uuid.UUID {
	return e.Id
}

func (e *FundsWithdrawn) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *FundsWithdrawn) GetVersion() int {
	return e.Version
}

func (e *FundsWithdrawn) SetVersion(number int) {
	e.Version = number
}

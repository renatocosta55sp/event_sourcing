package events

import "github.com/google/uuid"

type AccountMonthClosed struct {
	Id, AggregateId uuid.UUID
	Balance         float64
	Version         int
}

func (e *AccountMonthClosed) GetName() string {
	return "AccountingMonthClosed"
}

func (e *AccountMonthClosed) GetId() uuid.UUID {
	return e.Id
}

func (e *AccountMonthClosed) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *AccountMonthClosed) GetVersion() int {
	return e.Version
}

func (e *AccountMonthClosed) SetVersion(number int) {
	e.Version = number
}

package events

import "github.com/google/uuid"

type AccountMonthOpened struct {
	Id, AggregateId uuid.UUID
	Balance         float64
	Version         int
}

func (e *AccountMonthOpened) GetName() string {
	return "AccountingMonthOpened"
}

func (e *AccountMonthOpened) GetId() uuid.UUID {
	return e.Id
}

func (e *AccountMonthOpened) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *AccountMonthOpened) GetVersion() int {
	return e.Version
}

func (e *AccountMonthOpened) SetVersion(number int) {
	e.Version = number
}

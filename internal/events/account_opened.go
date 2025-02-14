package events

import "github.com/google/uuid"

type AccountOpened struct {
	Id, AccountId, OwnerId uuid.UUID
	InitialBalance         float64
	OpenedAt               string
	Version                int
}

func (e *AccountOpened) GetName() string {
	return "AccountOpened"
}

func (e *AccountOpened) GetId() uuid.UUID {
	return e.Id
}

func (e *AccountOpened) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *AccountOpened) GetVersion() int {
	return e.Version
}

func (e *AccountOpened) SetVersion(number int) {
	e.Version = number
}

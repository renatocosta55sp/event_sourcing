package events

import "github.com/google/uuid"

type DeviceRemoved struct {
	Id, AggregateId uuid.UUID
	Version         int
}

func (e *DeviceRemoved) GetName() string {
	return "DeviceRemoved"
}

func (e *DeviceRemoved) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *DeviceRemoved) GetId() uuid.UUID {
	return e.Id
}

func (e *DeviceRemoved) GetVersion() int {
	return e.Version
}

func (e *DeviceRemoved) SetVersion(number int) {
	e.Version = number
}

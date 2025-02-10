package events

import "github.com/google/uuid"

type DeviceAdded struct {
	Id, AggregateId uuid.UUID
	Name, Brand     string
	Version         int
}

func (e *DeviceAdded) GetName() string {
	return "DeviceAdded"
}

func (e *DeviceAdded) GetId() uuid.UUID {
	return e.Id
}

func (e *DeviceAdded) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *DeviceAdded) GetVersion() int {
	return e.Version
}

func (e *DeviceAdded) SetVersion(number int) {
	e.Version = number
}

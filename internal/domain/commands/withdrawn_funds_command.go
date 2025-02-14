package commands

import "github.com/google/uuid"

type WithdrawnFundsCommand struct {
	AggregateID uuid.UUID
	Amount      float64
}

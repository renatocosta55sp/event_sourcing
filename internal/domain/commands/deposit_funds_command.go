package commands

import "github.com/google/uuid"

type DepositFundsCommand struct {
	AggregateID uuid.UUID
	Amount      float64
}

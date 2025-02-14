package domain

import (
	"testing"

	"github.com/renatocosta55sp/event_sourcing/internal/domain/commands"
	"github.com/renatocosta55sp/modeling/domain"
	"github.com/stretchr/testify/assert"
)

func TestInvalidArguments(t *testing.T) {

	bankAccount := NewBankAccountAggregate([]domain.Event{})

	var tests = []struct {
		amount float64
		want   error
	}{
		{

			amount: -1,
			want:   ErrAmountMustBePositive,
		},
		{

			amount: 0,
			want:   ErrAmountMustBePositive,
		},
		{

			amount: -2,
			want:   ErrAmountMustBePositive,
		},
	}

	for _, test := range tests {
		_, err := bankAccount.Deposit(
			commands.DepositFundsCommand{
				AggregateID: bankAccount.AggregateID,
				Amount:      test.amount,
			},
		)
		assert.Equal(t, err, test.want, "Expected: %d - Got: %d", test.want, err)
	}

}

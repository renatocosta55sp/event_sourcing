package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetNextVal(ctx context.Context, db *pgxpool.Conn, sequenceName string) (int64, error) {
	var nextVal int64
	query := "SELECT nextval($1)"
	err := db.QueryRow(ctx, query, sequenceName).Scan(&nextVal)
	if err != nil {
		return 0, err
	}
	return nextVal, nil
}

type TransactionDb struct {
	Ctx  context.Context
	Conn *pgxpool.Conn
}

func (t *TransactionDb) Transaction(fn func() error) error {

	defer t.Conn.Release()

	tx, err := t.Conn.BeginTx(t.Ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %v", err)
	}

	defer func() {
		if err := tx.Rollback(t.Ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Printf("failed to rollback transaction: %v\n", err)
		}
	}()

	if err := fn(); err != nil {
		return err
	}

	err = tx.Commit(t.Ctx)
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

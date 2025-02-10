package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/renatocosta55sp/modeling/domain"
	"github.com/renatocosta55sp/modeling/eventstore"
	"github.com/renatocosta55sp/modeling/infra"
	"github.com/renatocosta55sp/modeling/infra/bus"
)

const SnapshotEventTableName = "snapshot_event_entry"

type PersistentSnapshotEventStore struct {
	Conn          *pgxpool.Pool
	eventRegistry *bus.EventRegistry
	DBSchema      string
}

func NewPersistentSnapshotEventStore(conn *pgxpool.Pool, eventRegistry *bus.EventRegistry, dbSchema string) eventstore.SnapshotStore {
	return &PersistentSnapshotEventStore{
		Conn:          conn,
		eventRegistry: eventRegistry,
		DBSchema:      dbSchema,
	}
}

func (p *PersistentSnapshotEventStore) WriteSnapshot(ctx context.Context, streamId string, event domain.Event, version int) error {

	entries := [][]any{}
	columns := []string{"sequence_number", "aggregate_identifier", "event_identifier", "time_stamp", "type", "meta_data", "payload"}

	metaDataSerialized, err := infra.Serialize(infra.EventMetadata{
		UserId: "System",
		Source: "WebApi",
	})

	if err != nil {
		return fmt.Errorf("failed to serialize metadata: %w", err)
	}

	payloadSerialized, err := infra.Serialize(event)
	if err != nil {
		return fmt.Errorf("failed to serialize payload: %w", err)
	}

	entries = append(entries, []any{version, streamId, event.GetId(), time.Now().Format("2006-01-02T15:04:05"), event.GetName(), metaDataSerialized, payloadSerialized})

	_, err = p.Conn.CopyFrom(
		ctx,
		pgx.Identifier{SnapshotEventTableName},
		columns,
		pgx.CopyFromRows(entries),
	)

	if err != nil {
		return fmt.Errorf("error copying into %s table: %w", SnapshotEventTableName, err)
	}

	return nil

}

func (p *PersistentSnapshotEventStore) ReadSnapshot(ctx context.Context, streamID string) (int, error) {

	var sequenceNumber int

	err := p.Conn.QueryRow(ctx, "SELECT sequence_number from snapshot_event_entry where aggregate_identifier = $1 ORDER BY sequence_number DESC LIMIT 1", streamID).Scan(&sequenceNumber)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil // No snapshot found, return 0 without error
		}
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return sequenceNumber, nil

}

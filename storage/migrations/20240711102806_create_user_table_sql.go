package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserTableSql, downCreateUserTableSql)
}

func upCreateUserTableSql(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE users (
			id UUID PRIMARY KEY,
			passport_number TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUserTableSql(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS tasks;
	`)
	if err != nil {
		return err
	}
	return nil
}

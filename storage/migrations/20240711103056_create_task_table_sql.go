package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTaskTableSql, downCreateTaskTableSql)
}

func upCreateTaskTableSql(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE tasks (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			start_time TIMESTAMP NOT NULL,
			end_time TIMESTAMP NOT NULL,
			total TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateTaskTableSql(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS users;
	`)
	if err != nil {
		return err
	}
	return nil
}

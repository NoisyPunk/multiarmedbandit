package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up001, Down001)
}

// Up001 up migration.
func Up001(tx *sql.Tx) error {
	query := `
			CREATE TABLE slots (
			id              uuid primary key,
            description text
            );
			CREATE TABLE banners (
			id              uuid primary key,
            description text
            );
			CREATE TABLE groups (
			id              uuid primary key,
            description text
            );
			`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// Down001 down migration.
func Down001(tx *sql.Tx) error {
	query := `
			DROP TABLE slots;
			DROP TABLE banners;
			DROP TABLE groups;
			`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

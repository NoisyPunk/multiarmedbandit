package storage

import (
	"database/sql"
	"fmt"

	calendarconfig "github.com/NoisyPunk/multiarmedbandit/internal/configs"
	// need for work with migrations.
	_ "github.com/NoisyPunk/multiarmedbandit/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func Migrate(config *calendarconfig.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DSN.Host, config.DSN.Port, config.DSN.User, config.DSN.Password, config.DSN.DBName, config.DSN.Ssl)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(db, "./"); err != nil {
		return err
	}
	return nil
}

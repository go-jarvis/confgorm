package migration

import (
	"errors"
)

var (
	ErrNotMigrator = errors.New("db driver does not a migrator")
)

type Migrator interface {
	Migrate() error
}

func Magrate(driver interface{}) error {
	migrator, ok := driver.(Migrator)
	if !ok {
		return ErrNotMigrator
	}

	return migrator.Migrate()
}

package mysql

import (
	"fmt"

	"github.com/go-jarvis/confgorm/migration"
	"github.com/sirupsen/logrus"
)

var _ migration.Migrator = (*MysqlDriver)(nil)

func (my *MysqlDriver) Migrate() error {

	// nerver do action when database target is not same
	if my.DbName != my.MigrationDB.Name() {
		return fmt.Errorf("dsn dbname(%s) != migrator dbname(%s), skip", my.DbName, my.MigrationDB.Name())
	}

	err := my.AutoMigrate(my.MigrationDB.Tables()...)
	if err != nil {
		return err
	}

	logrus.Infof("auto migrate success: db(%s)", my.DbName)
	return nil
}

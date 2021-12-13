package mysql

import (
	"fmt"

	"github.com/go-jarvis/confgorm/migration"
	"github.com/sirupsen/logrus"
)

var _ migration.Migrator = (*MysqlDriver)(nil)

func (my *MysqlDriver) Migrate() {

	// nerver do action when database target is not same
	if my.DbName != my.MigrationDB.Name() {
		fmt.Printf("%+v", my)
		logrus.Warnf("dsn dbname(%s) != migrator dbname(%s), skip", my.DbName, my.MigrationDB.Name())
		return
	}

	my.AutoMigrate(my.MigrationDB.Tables()...)
}

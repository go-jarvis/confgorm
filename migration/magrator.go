package migration

import "github.com/sirupsen/logrus"

type Migrator interface {
	Migrate()
}

func Magrate(driver interface{}) {
	migrator, ok := driver.(Migrator)
	if !ok {
		logrus.Warnf("db driver does not a migrator")
		return
	}

	migrator.Migrate()

}

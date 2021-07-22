package __example__

import (
	"testing"

	"github.com/goutilx/confgorm"
)

func Test_Sqlite(t *testing.T) {
	sqlite := confgorm.Sqlite{
		DBFile: "sqlite.db",
	}
	sqlite.Init()

	userProcessAll(t, sqlite.DB())
}

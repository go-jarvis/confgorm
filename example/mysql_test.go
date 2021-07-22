package __example__

import (
	"testing"

	"github.com/goutilx/confgorm"
)

func Test_Mysql(t *testing.T) {
	my := confgorm.Mysql{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "Mysql12345",
		Dbname:   "demo",
	}
	my.Init()

	userProcessAll(t, my.DB())
}

package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDriver struct {
	Host               string `env:""`
	Port               int    `env:""`
	User               string `env:""`
	Password           string `env:""`
	Dbname             string `env:""`
	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxIdleSeconds int

	*gorm.DB
}

func (my *MysqlDriver) SetDefaults() {
	if my.Port == 0 {
		my.Port = 3306
	}
	if my.Host == "" {
		my.Host = "127.0.0.1"
	}

	if my.MaxIdleConns == 0 {
		my.MaxIdleConns = 10
	}
	if my.MaxOpenConns == 0 {
		my.MaxOpenConns = 30
	}
	if my.ConnMaxIdleSeconds == 0 {
		my.ConnMaxIdleSeconds = 1800
	}
}

func (my *MysqlDriver) Init() {
	my.SetDefaults()
	if my.DB == nil {
		my.conn()
	}
}

// livenessProbe liveness checking
func (my *MysqlDriver) livenessProbe() {

}

// conn database connection
func (my *MysqlDriver) conn() error {

	_dsn_ := `%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`
	dsn := fmt.Sprintf(_dsn_,
		my.User, my.Password,
		my.Host, my.Port,
		my.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqldb, err := db.DB()
	if err != nil {
		return err
	}

	sqldb.SetConnMaxIdleTime(time.Duration(my.ConnMaxIdleSeconds) * time.Second)
	sqldb.SetMaxIdleConns(my.MaxIdleConns)
	sqldb.SetMaxOpenConns(my.MaxOpenConns)
	my.DB = db

	return nil
}

func (my *MysqlDriver) DBDriver() *gorm.DB {
	return my.DB
}

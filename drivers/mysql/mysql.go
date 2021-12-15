package mysql

import (
	"fmt"
	"time"

	"github.com/go-jarvis/confgorm/migration"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDriver struct {
	Host               string `env:""`
	Port               int    `env:""`
	User               string `env:""`
	Password           string `env:""`
	DbName             string `env:""`
	ConnectionOptions  string `env:""`
	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxIdleSeconds int

	// database to magrate
	MigrationDB *migration.Database `env:"-"`

	*gorm.DB `env:"-"`
}

func (my *MysqlDriver) SetDefaults() {
	if my.Port == 0 {
		my.Port = 3306
	}
	if my.Host == "" {
		my.Host = "127.0.0.1"
	}

	if my.ConnectionOptions == "" {
		my.ConnectionOptions = "charset=utf8mb4&parseTime=True&loc=Local"
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
		_ = my.conn()
	}

	go my.livenessChecking()
}

func (my *MysqlDriver) ping() error {
	sqldb, err := my.DB.DB()
	if err != nil {
		return err
	}

	return sqldb.Ping()
}

func (my *MysqlDriver) retry(counter int) (err error) {

	// max retry interval 30s
	if counter > 6 {
		counter = 6
	}
	t := time.Duration(counter) * 5 * time.Second
	time.Sleep(t)

	err = my.conn()
	if err == nil {
		return nil
	}

	return err
}

// livenessChecking liveness checking
func (my *MysqlDriver) livenessChecking() {

	for {
		// liveness checking every 60s
		err := my.ping()
		if err == nil {
			time.Sleep(60 * time.Second)
			continue
		}

		logrus.Errorf("db ping failed: %v", err)
		// retry
		counter := 0
		for {
			err := my.retry(counter)
			if err != nil {
				logrus.Errorf("db retried to connect %d times", counter)

				counter += 1
				continue
			}

			break
		}
	}
}

// conn database connection
func (my *MysqlDriver) conn() error {

	_dsn_ := `%s:%s@tcp(%s:%d)/%s?%s`
	dsn := fmt.Sprintf(_dsn_,
		my.User, my.Password,
		my.Host, my.Port,
		my.DbName,

		my.ConnectionOptions)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqldb, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqldb.Ping(); err != nil {
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

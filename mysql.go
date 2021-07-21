package confmysql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host               string `env:""`
	Port               int    `env:""`
	User               string `env:""`
	Password           string `env:""`
	Dbname             string `env:""`
	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxIdleSeconds int

	db *gorm.DB
}

func (my *Mysql) SetDefaults() {
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
		my.MaxOpenConns = 20
	}
	if my.ConnMaxIdleSeconds == 0 {
		my.ConnMaxIdleSeconds = 30
	}
}

func (my *Mysql) Init() {
	my.SetDefaults()
	if my.db == nil {
		my.initial()
	}
}

func (my *Mysql) initial() {

	_dsn_ := `%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`
	dsn := fmt.Sprintf(_dsn_,
		my.User, my.Password,
		my.Host, my.Port,
		my.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqldb, err := db.DB()
	if err != nil {
		sqldb.SetConnMaxIdleTime(time.Duration(my.ConnMaxIdleSeconds) * time.Second)
		sqldb.SetMaxIdleConns(my.MaxIdleConns)
		sqldb.SetMaxOpenConns(my.MaxOpenConns)
	}

	my.db = db
}

func (my *Mysql) DB() *gorm.DB {
	return my.db
}
